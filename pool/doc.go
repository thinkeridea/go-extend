// MIT License
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

// Package pool 创建一些常用的 pool， 底层依托 sync.pool
//
// 在项目开发中会大量使用 sync.Pool, 有很多 pool 是可以共享的，最常见的就是 bytes.Buffer。
// 大量的开源库中在使用 bytes.Buffer 的 sync.Pool, 每个项目单独维护自己的 sync.Pool，
// 使得很多资源没有被合理的利用， 造成大量的浪费，这个包收集常见的 sync.Pool 实例，使其可以在不同项目和包中可以共享。
package pool
