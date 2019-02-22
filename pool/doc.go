// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package pool 创建一些常用的 pool， 底层依托 sync.pool
//
// 在项目开发中会大量使用 sync.Pool, 有很多 pool 是可以共享的，最常见的就是 bytes.Buffer。
// 大量的开源库中在使用 bytes.Buffer 的 sync.Pool, 每个项目单独维护自己的 sync.Pool，
// 使得很多资源没有被合理的利用， 造成大量的浪费，这个包收集常见的 sync.Pool 实例，使其可以在不同项目和包中可以共享。
package pool
