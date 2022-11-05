// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file is adapted from:
// https://github.com/hajimehoshi/ebiten/blob/main/examples/shader/lighting.go

package main

// var Time float
// var Cursor vec2
// var ScreenSize vec2
var TorchPos vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	lightpos := vec3(TorchPos, 50)
	lightdir := normalize(lightpos - position.xyz)
	normal := normalize(imageSrc1UnsafeAt(texCoord) - 0.5)
	const ambient = 0.25
	diffuse := 0.75 * max(0.0, dot(normal.xyz, lightdir))
	return imageSrc0UnsafeAt(texCoord) * (ambient + diffuse)
}

// void main(){
// 	vec2 st = gl_FragCoord.xy/u_resolution;
//     float pct = 0.0;

//     // a. The DISTANCE from the pixel to the center
//    pct = 1.0-distance(vec2(0.500,0.550),st);

//     vec3 color = vec3(smoothstep(0.7,0.9,pct));

// 	gl_FragColor = vec4( color, 1.0 );
// }