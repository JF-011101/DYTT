/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:25:02
 * @FilePath: \DYTT\cmd\api\rpc\init.go
 * @Description: Initialize RPC client based on configuration information
 */

package rpc

import "github.com/jf-011101/dytt/pkg/ttviper"

// InitRPC init rpc client
func InitRPC(Config *ttviper.Config) {
	UserConfig := ttviper.ConfigInit("TIKTOK_USER", "userConfig")
	initUserRpc(&UserConfig)

	FeedConfig := ttviper.ConfigInit("TIKTOK_FEED", "feedConfig")
	initFeedRpc(&FeedConfig)

	PublishConfig := ttviper.ConfigInit("TIKTOK_PUBLISH", "publishConfig")
	initPublishRpc(&PublishConfig)

	FavoriteConfig := ttviper.ConfigInit("TIKTOK_FAVORITE", "favoriteConfig")
	initFavoriteRpc(&FavoriteConfig)

	CommentConfig := ttviper.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	initCommentRpc(&CommentConfig)

	RelationConfig := ttviper.ConfigInit("TIKTOK_RELATION", "relationConfig")
	initRelationRpc(&RelationConfig)
}
