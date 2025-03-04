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

type DataTransportSize uint8

type IDataTransportSize interface {
    SizeInBits() bool
    Serialize(io utils.WriteBuffer) error
}

const(
    DataTransportSize_NULL DataTransportSize = 0x00
    DataTransportSize_BIT DataTransportSize = 0x03
    DataTransportSize_BYTE_WORD_DWORD DataTransportSize = 0x04
    DataTransportSize_INTEGER DataTransportSize = 0x05
    DataTransportSize_DINTEGER DataTransportSize = 0x06
    DataTransportSize_REAL DataTransportSize = 0x07
    DataTransportSize_OCTET_STRING DataTransportSize = 0x09
)


func (e DataTransportSize) SizeInBits() bool {
    switch e  {
        case 0x00: { /* '0x00' */
            return false
        }
        case 0x03: { /* '0x03' */
            return true
        }
        case 0x04: { /* '0x04' */
            return true
        }
        case 0x05: { /* '0x05' */
            return true
        }
        case 0x06: { /* '0x06' */
            return false
        }
        case 0x07: { /* '0x07' */
            return false
        }
        case 0x09: { /* '0x09' */
            return false
        }
        default: {
            return false
        }
    }
}
func DataTransportSizeValueOf(value uint8) DataTransportSize {
    switch value {
        case 0x00:
            return DataTransportSize_NULL
        case 0x03:
            return DataTransportSize_BIT
        case 0x04:
            return DataTransportSize_BYTE_WORD_DWORD
        case 0x05:
            return DataTransportSize_INTEGER
        case 0x06:
            return DataTransportSize_DINTEGER
        case 0x07:
            return DataTransportSize_REAL
        case 0x09:
            return DataTransportSize_OCTET_STRING
    }
    return 0
}

func CastDataTransportSize(structType interface{}) DataTransportSize {
    castFunc := func(typ interface{}) DataTransportSize {
        if sDataTransportSize, ok := typ.(DataTransportSize); ok {
            return sDataTransportSize
        }
        return 0
    }
    return castFunc(structType)
}

func (m DataTransportSize) LengthInBits() uint16 {
    return 8
}

func (m DataTransportSize) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DataTransportSizeParse(io *utils.ReadBuffer) (DataTransportSize, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return DataTransportSizeValueOf(val), nil
}

func (e DataTransportSize) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}

func (e DataTransportSize) String() string {
    switch e {
    case DataTransportSize_NULL:
        return "NULL"
    case DataTransportSize_BIT:
        return "BIT"
    case DataTransportSize_BYTE_WORD_DWORD:
        return "BYTE_WORD_DWORD"
    case DataTransportSize_INTEGER:
        return "INTEGER"
    case DataTransportSize_DINTEGER:
        return "DINTEGER"
    case DataTransportSize_REAL:
        return "REAL"
    case DataTransportSize_OCTET_STRING:
        return "OCTET_STRING"
    }
    return ""
}
