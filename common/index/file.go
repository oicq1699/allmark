// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"github.com/andreaskoch/allmark2/common/logger"
	"github.com/andreaskoch/allmark2/common/route"
	"github.com/andreaskoch/allmark2/model"
)

func CreateFileIndex(logger logger.Logger) *FileIndex {
	return &FileIndex{
		logger: logger,
		items:  make(map[route.Route]*model.File),
	}
}

type FileIndex struct {
	logger logger.Logger
	items  map[route.Route]*model.File
}

func (index *FileIndex) IsMatch(route route.Route) (file *model.File, isMatch bool) {
	file, isMatch = index.items[route]
	return
}

func (index *FileIndex) Routes() []route.Route {
	routes := make([]route.Route, 0)
	for route, _ := range index.items {
		routes = append(routes, route)
	}
	return routes
}

func (index *FileIndex) Add(file *model.File) {
	index.logger.Debug("Adding file %q to index", file)
	index.items[*file.Route()] = file
}