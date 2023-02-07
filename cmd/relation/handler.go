package main

import (
	"context"
	relation "douyin/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollow(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollower(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
