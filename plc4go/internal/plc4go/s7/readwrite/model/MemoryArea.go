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

type MemoryArea uint8

type IMemoryArea interface {
    ShortName() string
    Serialize(io utils.WriteBuffer) error
}

const(
    MemoryArea_COUNTERS MemoryArea = 0x1C
    MemoryArea_TIMERS MemoryArea = 0x1D
    MemoryArea_DIRECT_PERIPHERAL_ACCESS MemoryArea = 0x80
    MemoryArea_INPUTS MemoryArea = 0x81
    MemoryArea_OUTPUTS MemoryArea = 0x82
    MemoryArea_FLAGS_MARKERS MemoryArea = 0x83
    MemoryArea_DATA_BLOCKS MemoryArea = 0x84
    MemoryArea_INSTANCE_DATA_BLOCKS MemoryArea = 0x85
    MemoryArea_LOCAL_DATA MemoryArea = 0x86
)


func (e MemoryArea) ShortName() string {
    switch e  {
        case 0x1C: { /* '0x1C' */
            return "C"
        }
        case 0x1D: { /* '0x1D' */
            return "T"
        }
        case 0x80: { /* '0x80' */
            return "D"
        }
        case 0x81: { /* '0x81' */
            return "I"
        }
        case 0x82: { /* '0x82' */
            return "Q"
        }
        case 0x83: { /* '0x83' */
            return "M"
        }
        case 0x84: { /* '0x84' */
            return "DB"
        }
        case 0x85: { /* '0x85' */
            return "DBI"
        }
        case 0x86: { /* '0x86' */
            return "LD"
        }
        default: {
            return ""
        }
    }
}
func MemoryAreaValueOf(value uint8) MemoryArea {
    switch value {
        case 0x1C:
            return MemoryArea_COUNTERS
        case 0x1D:
            return MemoryArea_TIMERS
        case 0x80:
            return MemoryArea_DIRECT_PERIPHERAL_ACCESS
        case 0x81:
            return MemoryArea_INPUTS
        case 0x82:
            return MemoryArea_OUTPUTS
        case 0x83:
            return MemoryArea_FLAGS_MARKERS
        case 0x84:
            return MemoryArea_DATA_BLOCKS
        case 0x85:
            return MemoryArea_INSTANCE_DATA_BLOCKS
        case 0x86:
            return MemoryArea_LOCAL_DATA
    }
    return 0
}

func CastMemoryArea(structType interface{}) MemoryArea {
    castFunc := func(typ interface{}) MemoryArea {
        if sMemoryArea, ok := typ.(MemoryArea); ok {
            return sMemoryArea
        }
        return 0
    }
    return castFunc(structType)
}

func (m MemoryArea) LengthInBits() uint16 {
    return 8
}

func (m MemoryArea) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func MemoryAreaParse(io *utils.ReadBuffer) (MemoryArea, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return MemoryAreaValueOf(val), nil
}

func (e MemoryArea) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}

func (e MemoryArea) String() string {
    switch e {
    case MemoryArea_COUNTERS:
        return "COUNTERS"
    case MemoryArea_TIMERS:
        return "TIMERS"
    case MemoryArea_DIRECT_PERIPHERAL_ACCESS:
        return "DIRECT_PERIPHERAL_ACCESS"
    case MemoryArea_INPUTS:
        return "INPUTS"
    case MemoryArea_OUTPUTS:
        return "OUTPUTS"
    case MemoryArea_FLAGS_MARKERS:
        return "FLAGS_MARKERS"
    case MemoryArea_DATA_BLOCKS:
        return "DATA_BLOCKS"
    case MemoryArea_INSTANCE_DATA_BLOCKS:
        return "INSTANCE_DATA_BLOCKS"
    case MemoryArea_LOCAL_DATA:
        return "LOCAL_DATA"
    }
    return ""
}
