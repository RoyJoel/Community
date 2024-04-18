package impl

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/zhangjiacheng-iHealth/IHCommunity/package/cache"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/config"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/middleware"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/model"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/utils"
	"github.com/zhangjiacheng-iHealth/IHCommunity/package/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

type IHCommunityDaoImpl struct {
	db    *gorm.DB
	cache *cache.IHCommunityCacheDAOImpl
	sess *IHCommunityAWS.sess
}

func NewIHCommunityDaoImpl() *IHCommunityDaoImpl {
	return &IHCommunityDaoImpl{db: config.DB, cache: cache.NewIHCommunityCacheDAOImpl()}
}

func (impl *IHCommunityDaoImpl) GetCommentsForUser(ctx context.Context, userID int64, postID int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := impl.db.Where("UserId = ? AND PostId = ?", userID, postID).Find(&comments).Error; err != nil {
			return nil, err
	}
	return comments, nil
}

func (impl *IHCommunityDaoImpl) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	if err := impl.db.Find(&posts).Error; err != nil {
			return nil, err
	}
	return posts, nil
}

func (impl *IHCommunityDaoImpl) GetVersionLogsByPostID(ctx context.Context, postID int64) []model.VersionLog {
	var versionlogs []model.VersionLog
	impl.db.Where("PostId = ?", postID).Find(&versionlogs)
	return versionlogs
}

func (impl *IHCommunityDaoImpl) GetProposalInfo(ctx context.Context, proposalID int64) model.Proposal {
	var proposal model.Proposal
	impl.db.First(&proposal, proposalID)
	return proposal
}

func (impl *IHCommunityDaoImpl) GetPlansByProposalID(ctx context.Context, proposalID int64) []model.Plan {
	var plans []model.Plan
	impl.db.Where("ProposalId = ?", proposalID).Find(&plans)
	return plans
}

func (impl *IHCommunityDaoImpl) GetPrivilegesByProposalID(ctx context.Context, proposalID int64) []model.Privileges {
	var privileges []model.Privileges
	impl.db.Where("ProposalId = ?", proposalID).Find(&privileges)
	return privileges
}

func (impl *IHCommunityDaoImpl) DeleteComment(ctx context.Context, commentID int64) error {
	if err := impl.db.Delete(&model.Comment{}, commentID).Error; err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) DeletePost(ctx context.Context, postID int64) error {
	if err := impl.db.Transaction(func(tx *gorm.DB) error {
			// 删除帖子对应的所有版本日志
			if err := tx.Delete(&model.VersionLog{}, "post_id = ?", postID).Error; err != nil {
					return err
			}
			// 删除帖子
			if err := tx.Delete(&model.Post{}, postID).Error; err != nil {
					return err
			}
			return nil
	}); err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) UpdateComment(ctx context.Context, comment model.Comment) error {
	if err := impl.db.Save(&comment).Error; err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) UpdatePost(ctx context.Context, post model.Post) error {
	if err := impl.db.Transaction(func(tx *gorm.DB) error {
			// 更新帖子信息
			if err := tx.Save(&post).Error; err != nil {
					return err
			}
			// 更新帖子对应的所有版本日志的相关信息
			if err := tx.Model(&model.VersionLog{}).Where("post_id = ?", post.Id).Update(map[string]interface{}{
					"Title":   post.Title,
					"Content": post.Content,
			}).Error; err != nil {
					return err
			}
			return nil
	}); err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) AddUser(ctx context.Context, user model.User) error {
	if err := impl.db.Create(&user).Error; err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) AddPost(ctx context.Context, post model.Post) error {
	if err := impl.db.Transaction(func(tx *gorm.DB) error {
			// 添加帖子
			if err := tx.Create(&post).Error; err != nil {
					return err
			}
			// 如果需要，添加帖子对应的版本日志
			if post.VersionLogs != nil && len(post.VersionLogs) > 0 {
					for _, log := range post.VersionLogs {
							log.PostId = post.Id
							if err := tx.Create(&log).Error; err != nil {
									return err
							}
					}
			}
			return nil
	}); err != nil {
			return err
	}
	return nil
}

func (impl *IHCommunityDaoImpl) SearchUser(ctx context.Context, userID int64) (model.User, error) {
	var user model.User
	if err := impl.db.First(&user, userID).Error; err != nil {
			return model.User{}, err
	}
	return user, nil
}

	func (impl *IHCommunityDaoImpl) UploadImage(ctx context.Context, file multipart.File, filename string) error {
   // 读取文件内容
	 fileBytes, err := ioutil.ReadAll(file)
	 if err != nil {
			 return err
	 }

	 // 创建 S3 客户端
	 svc := s3.New(impl.sess)

	 // 设置 S3 存储桶名称和图片键
	 bucket := "zhangjiacheng"
	 key := filename

	 // 上传图片到 S3 存储桶
	 _, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
			 Bucket: aws.String(bucket),
			 Key:    aws.String(key),
			 Body:   bytes.NewReader(fileBytes),
	 })
	 if err != nil {
			 return err
	 }

	 return nil
}


func (awsS3 AwsS3) DeleteObject(bucket string, path string) (bool, error) {
	obj := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}
	res, err := awsS3.client.DeleteObject(context.TODO(), obj)
	if err != nil {
		return false, err
	}
	return *res.DeleteMarker, err
}