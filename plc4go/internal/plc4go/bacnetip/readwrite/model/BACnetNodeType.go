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

type BACnetNodeType uint8

type IBACnetNodeType interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    BACnetNodeType_UNKNOWN BACnetNodeType = 0x00
    BACnetNodeType_SYSTEM BACnetNodeType = 0x01
    BACnetNodeType_NETWORK BACnetNodeType = 0x02
    BACnetNodeType_DEVICE BACnetNodeType = 0x03
    BACnetNodeType_ORGANIZATIONAL BACnetNodeType = 0x04
    BACnetNodeType_AREA BACnetNodeType = 0x05
    BACnetNodeType_EQUIPMENT BACnetNodeType = 0x06
    BACnetNodeType_POINT BACnetNodeType = 0x07
    BACnetNodeType_COLLECTION BACnetNodeType = 0x08
    BACnetNodeType_PROPERTY BACnetNodeType = 0x09
    BACnetNodeType_FUNCTIONAL BACnetNodeType = 0x0A
    BACnetNodeType_OTHER BACnetNodeType = 0x0B
    BACnetNodeType_SUBSYSTEM BACnetNodeType = 0x0C
    BACnetNodeType_BUILDING BACnetNodeType = 0x0D
    BACnetNodeType_FLOOR BACnetNodeType = 0x0E
    BACnetNodeType_SECTION BACnetNodeType = 0x0F
    BACnetNodeType_MODULE BACnetNodeType = 0x10
    BACnetNodeType_TREE BACnetNodeType = 0x11
    BACnetNodeType_MEMBER BACnetNodeType = 0x12
    BACnetNodeType_PROTOCOL BACnetNodeType = 0x13
    BACnetNodeType_ROOM BACnetNodeType = 0x14
    BACnetNodeType_ZONE BACnetNodeType = 0x15
)

func BACnetNodeTypeValueOf(value uint8) BACnetNodeType {
    switch value {
        case 0x00:
            return BACnetNodeType_UNKNOWN
        case 0x01:
            return BACnetNodeType_SYSTEM
        case 0x02:
            return BACnetNodeType_NETWORK
        case 0x03:
            return BACnetNodeType_DEVICE
        case 0x04:
            return BACnetNodeType_ORGANIZATIONAL
        case 0x05:
            return BACnetNodeType_AREA
        case 0x06:
            return BACnetNodeType_EQUIPMENT
        case 0x07:
            return BACnetNodeType_POINT
        case 0x08:
            return BACnetNodeType_COLLECTION
        case 0x09:
            return BACnetNodeType_PROPERTY
        case 0x0A:
            return BACnetNodeType_FUNCTIONAL
        case 0x0B:
            return BACnetNodeType_OTHER
        case 0x0C:
            return BACnetNodeType_SUBSYSTEM
        case 0x0D:
            return BACnetNodeType_BUILDING
        case 0x0E:
            return BACnetNodeType_FLOOR
        case 0x0F:
            return BACnetNodeType_SECTION
        case 0x10:
            return BACnetNodeType_MODULE
        case 0x11:
            return BACnetNodeType_TREE
        case 0x12:
            return BACnetNodeType_MEMBER
        case 0x13:
            return BACnetNodeType_PROTOCOL
        case 0x14:
            return BACnetNodeType_ROOM
        case 0x15:
            return BACnetNodeType_ZONE
    }
    return 0
}

func CastBACnetNodeType(structType interface{}) BACnetNodeType {
    castFunc := func(typ interface{}) BACnetNodeType {
        if sBACnetNodeType, ok := typ.(BACnetNodeType); ok {
            return sBACnetNodeType
        }
        return 0
    }
    return castFunc(structType)
}

func (m BACnetNodeType) LengthInBits() uint16 {
    return 8
}

func (m BACnetNodeType) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetNodeTypeParse(io *utils.ReadBuffer) (BACnetNodeType, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return BACnetNodeTypeValueOf(val), nil
}

func (e BACnetNodeType) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}

func (e BACnetNodeType) String() string {
    switch e {
    case BACnetNodeType_UNKNOWN:
        return "UNKNOWN"
    case BACnetNodeType_SYSTEM:
        return "SYSTEM"
    case BACnetNodeType_NETWORK:
        return "NETWORK"
    case BACnetNodeType_DEVICE:
        return "DEVICE"
    case BACnetNodeType_ORGANIZATIONAL:
        return "ORGANIZATIONAL"
    case BACnetNodeType_AREA:
        return "AREA"
    case BACnetNodeType_EQUIPMENT:
        return "EQUIPMENT"
    case BACnetNodeType_POINT:
        return "POINT"
    case BACnetNodeType_COLLECTION:
        return "COLLECTION"
    case BACnetNodeType_PROPERTY:
        return "PROPERTY"
    case BACnetNodeType_FUNCTIONAL:
        return "FUNCTIONAL"
    case BACnetNodeType_OTHER:
        return "OTHER"
    case BACnetNodeType_SUBSYSTEM:
        return "SUBSYSTEM"
    case BACnetNodeType_BUILDING:
        return "BUILDING"
    case BACnetNodeType_FLOOR:
        return "FLOOR"
    case BACnetNodeType_SECTION:
        return "SECTION"
    case BACnetNodeType_MODULE:
        return "MODULE"
    case BACnetNodeType_TREE:
        return "TREE"
    case BACnetNodeType_MEMBER:
        return "MEMBER"
    case BACnetNodeType_PROTOCOL:
        return "PROTOCOL"
    case BACnetNodeType_ROOM:
        return "ROOM"
    case BACnetNodeType_ZONE:
        return "ZONE"
    }
    return ""
}
