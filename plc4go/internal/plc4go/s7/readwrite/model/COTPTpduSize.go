//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type COTPTpduSize int8

type ICOTPTpduSize interface {
    SizeInBytes() uint16
    Serialize(io utils.WriteBuffer) error
}

const(
    COTPTpduSize_SIZE_128 COTPTpduSize = 0x07
    COTPTpduSize_SIZE_256 COTPTpduSize = 0x08
    COTPTpduSize_SIZE_512 COTPTpduSize = 0x09
    COTPTpduSize_SIZE_1024 COTPTpduSize = 0x0a
    COTPTpduSize_SIZE_2048 COTPTpduSize = 0x0b
    COTPTpduSize_SIZE_4096 COTPTpduSize = 0x0c
    COTPTpduSize_SIZE_8192 COTPTpduSize = 0x0d
)


func (e COTPTpduSize) SizeInBytes() uint16 {
    switch e  {
        case 0x07: { /* '0x07' */
            return 128
        }
        case 0x08: { /* '0x08' */
            return 256
        }
        case 0x09: { /* '0x09' */
            return 512
        }
        case 0x0a: { /* '0x0a' */
            return 1024
        }
        case 0x0b: { /* '0x0b' */
            return 2048
        }
        case 0x0c: { /* '0x0c' */
            return 4096
        }
        case 0x0d: { /* '0x0d' */
            return 8192
        }
        default: {
            return 0
        }
    }
}
func COTPTpduSizeValueOf(value int8) COTPTpduSize {
    switch value {
        case 0x07:
            return COTPTpduSize_SIZE_128
        case 0x08:
            return COTPTpduSize_SIZE_256
        case 0x09:
            return COTPTpduSize_SIZE_512
        case 0x0a:
            return COTPTpduSize_SIZE_1024
        case 0x0b:
            return COTPTpduSize_SIZE_2048
        case 0x0c:
            return COTPTpduSize_SIZE_4096
        case 0x0d:
            return COTPTpduSize_SIZE_8192
    }
    return 0
}

func CastCOTPTpduSize(structType interface{}) COTPTpduSize {
    castFunc := func(typ interface{}) COTPTpduSize {
        if sCOTPTpduSize, ok := typ.(COTPTpduSize); ok {
            return sCOTPTpduSize
        }
        return 0
    }
    return castFunc(structType)
}

func (m COTPTpduSize) LengthInBits() uint16 {
    return 8
}

func (m COTPTpduSize) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPTpduSizeParse(io *utils.ReadBuffer) (COTPTpduSize, error) {
    val, err := io.ReadInt8(8)
    if err != nil {
        return 0, nil
    }
    return COTPTpduSizeValueOf(val), nil
}

func (e COTPTpduSize) Serialize(io utils.WriteBuffer) error {
    err := io.WriteInt8(8, int8(e))
    return err
}

func (e COTPTpduSize) String() string {
    switch e {
    case COTPTpduSize_SIZE_128:
        return "SIZE_128"
    case COTPTpduSize_SIZE_256:
        return "SIZE_256"
    case COTPTpduSize_SIZE_512:
        return "SIZE_512"
    case COTPTpduSize_SIZE_1024:
        return "SIZE_1024"
    case COTPTpduSize_SIZE_2048:
        return "SIZE_2048"
    case COTPTpduSize_SIZE_4096:
        return "SIZE_4096"
    case COTPTpduSize_SIZE_8192:
        return "SIZE_8192"
    }
    return ""
}
