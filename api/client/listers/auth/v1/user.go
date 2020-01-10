/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/auth/v1"
)

// UserLister helps list Users.
type UserLister interface {
	// List lists all Users in the indexer.
	List(selector labels.Selector) (ret []*v1.User, err error)
	// Get retrieves the User from the index for a given name.
	Get(name string) (*v1.User, error)
	UserListerExpansion
}

// userLister implements the UserLister interface.
type userLister struct {
	indexer cache.Indexer
}

// NewUserLister returns a new UserLister.
func NewUserLister(indexer cache.Indexer) UserLister {
	return &userLister{indexer: indexer}
}

// List lists all Users in the indexer.
func (s *userLister) List(selector labels.Selector) (ret []*v1.User, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.User))
	})
	return ret, err
}

// Get retrieves the User from the index for a given name.
func (s *userLister) Get(name string) (*v1.User, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("user"), name)
	}
	return obj.(*v1.User), nil
}
