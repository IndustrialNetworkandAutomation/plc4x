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

type BACnetObjectType uint16

type IBACnetObjectType interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    BACnetObjectType_ANALOG_INPUT BACnetObjectType = 0x000
    BACnetObjectType_ANALOG_OUTPUT BACnetObjectType = 0x001
    BACnetObjectType_ANALOG_VALUE BACnetObjectType = 0x002
    BACnetObjectType_BINARY_INPUT BACnetObjectType = 0x003
    BACnetObjectType_BINARY_OUTPUT BACnetObjectType = 0x004
    BACnetObjectType_BINARY_VALUE BACnetObjectType = 0x005
    BACnetObjectType_MULTISTATE_INPUT BACnetObjectType = 0x00D
    BACnetObjectType_MULTISTATE_OUTPUT BACnetObjectType = 0x00E
    BACnetObjectType_MULTISTATE_VALUE BACnetObjectType = 0x013
    BACnetObjectType_SCHEDULE BACnetObjectType = 0x011
    BACnetObjectType_DEVICE BACnetObjectType = 0x008
    BACnetObjectType_CALENDAR BACnetObjectType = 0x006
    BACnetObjectType_COMMAND BACnetObjectType = 0x007
    BACnetObjectType_EVENT_ENROLLMENT BACnetObjectType = 0x009
    BACnetObjectType_FILE BACnetObjectType = 0x00A
    BACnetObjectType_GROUP BACnetObjectType = 0x00B
    BACnetObjectType_LOOP BACnetObjectType = 0x00C
    BACnetObjectType_NOTIFICATION_CLASS BACnetObjectType = 0x00F
    BACnetObjectType_PROGRAM BACnetObjectType = 0x010
    BACnetObjectType_AVERAGING BACnetObjectType = 0x012
    BACnetObjectType_TREND_LOG BACnetObjectType = 0x014
    BACnetObjectType_LIFE_SAFETY_POINT BACnetObjectType = 0x015
    BACnetObjectType_LIFE_SAFETY_ZONE BACnetObjectType = 0x016
    BACnetObjectType_ACCUMULATOR BACnetObjectType = 0x017
    BACnetObjectType_PULSE_CONVERTER BACnetObjectType = 0x018
    BACnetObjectType_EVENT_LOG BACnetObjectType = 0x019
    BACnetObjectType_GLOBAL_GROUP BACnetObjectType = 0x01A
    BACnetObjectType_TREND_LOG_MULTIPLE BACnetObjectType = 0x01B
    BACnetObjectType_LOAD_CONTROL BACnetObjectType = 0x01C
    BACnetObjectType_STRUCTURED_VIEW BACnetObjectType = 0x01D
    BACnetObjectType_ACCESS_DOOR BACnetObjectType = 0x01E
    BACnetObjectType_TIMER BACnetObjectType = 0x01F
    BACnetObjectType_ACCESS_CREDENTIAL BACnetObjectType = 0x020
    BACnetObjectType_ACCESS_POINT BACnetObjectType = 0x021
    BACnetObjectType_ACCESS_RIGHTS BACnetObjectType = 0x022
    BACnetObjectType_ACCESS_USER BACnetObjectType = 0x023
    BACnetObjectType_ACCESS_ZONE BACnetObjectType = 0x024
    BACnetObjectType_CREDENTIAL_DATA_INPUT BACnetObjectType = 0x025
    BACnetObjectType_NETWORK_SECURITY BACnetObjectType = 0x026
    BACnetObjectType_BITSTRING_VALUE BACnetObjectType = 0x027
    BACnetObjectType_CHARACTERSTRING_VALUE BACnetObjectType = 0x028
    BACnetObjectType_DATEPATTERN_VALUE BACnetObjectType = 0x029
    BACnetObjectType_DATE_VALUE BACnetObjectType = 0x02A
    BACnetObjectType_DATETIMEPATTERN_VALUE BACnetObjectType = 0x02B
    BACnetObjectType_DATETIME_VALUE BACnetObjectType = 0x02C
    BACnetObjectType_INTEGER_VALUE BACnetObjectType = 0x02D
    BACnetObjectType_LARGE_ANALOG_VALUE BACnetObjectType = 0x02E
    BACnetObjectType_OCTETSTRING_VALUE BACnetObjectType = 0x02F
    BACnetObjectType_POSITIVE_INTEGER_VALUE BACnetObjectType = 0x030
    BACnetObjectType_TIMEPATTERN_VALUE BACnetObjectType = 0x031
    BACnetObjectType_TIME_VALUE BACnetObjectType = 0x032
    BACnetObjectType_NOTIFICATION_FORWARDER BACnetObjectType = 0x033
    BACnetObjectType_ALERT_ENROLLMENT BACnetObjectType = 0x034
    BACnetObjectType_CHANNEL BACnetObjectType = 0x035
    BACnetObjectType_LIGHTING_OUTPUT BACnetObjectType = 0x036
    BACnetObjectType_BINARY_LIGHTING_OUTPUT BACnetObjectType = 0x037
    BACnetObjectType_NETWORK_PORT BACnetObjectType = 0x038
    BACnetObjectType_ELEVATOR_GROUP BACnetObjectType = 0x039
    BACnetObjectType_ESCALATOR BACnetObjectType = 0x03A
)

func BACnetObjectTypeValueOf(value uint16) BACnetObjectType {
    switch value {
        case 0x000:
            return BACnetObjectType_ANALOG_INPUT
        case 0x001:
            return BACnetObjectType_ANALOG_OUTPUT
        case 0x002:
            return BACnetObjectType_ANALOG_VALUE
        case 0x003:
            return BACnetObjectType_BINARY_INPUT
        case 0x004:
            return BACnetObjectType_BINARY_OUTPUT
        case 0x005:
            return BACnetObjectType_BINARY_VALUE
        case 0x006:
            return BACnetObjectType_CALENDAR
        case 0x007:
            return BACnetObjectType_COMMAND
        case 0x008:
            return BACnetObjectType_DEVICE
        case 0x009:
            return BACnetObjectType_EVENT_ENROLLMENT
        case 0x00A:
            return BACnetObjectType_FILE
        case 0x00B:
            return BACnetObjectType_GROUP
        case 0x00C:
            return BACnetObjectType_LOOP
        case 0x00D:
            return BACnetObjectType_MULTISTATE_INPUT
        case 0x00E:
            return BACnetObjectType_MULTISTATE_OUTPUT
        case 0x00F:
            return BACnetObjectType_NOTIFICATION_CLASS
        case 0x010:
            return BACnetObjectType_PROGRAM
        case 0x011:
            return BACnetObjectType_SCHEDULE
        case 0x012:
            return BACnetObjectType_AVERAGING
        case 0x013:
            return BACnetObjectType_MULTISTATE_VALUE
        case 0x014:
            return BACnetObjectType_TREND_LOG
        case 0x015:
            return BACnetObjectType_LIFE_SAFETY_POINT
        case 0x016:
            return BACnetObjectType_LIFE_SAFETY_ZONE
        case 0x017:
            return BACnetObjectType_ACCUMULATOR
        case 0x018:
            return BACnetObjectType_PULSE_CONVERTER
        case 0x019:
            return BACnetObjectType_EVENT_LOG
        case 0x01A:
            return BACnetObjectType_GLOBAL_GROUP
        case 0x01B:
            return BACnetObjectType_TREND_LOG_MULTIPLE
        case 0x01C:
            return BACnetObjectType_LOAD_CONTROL
        case 0x01D:
            return BACnetObjectType_STRUCTURED_VIEW
        case 0x01E:
            return BACnetObjectType_ACCESS_DOOR
        case 0x01F:
            return BACnetObjectType_TIMER
        case 0x020:
            return BACnetObjectType_ACCESS_CREDENTIAL
        case 0x021:
            return BACnetObjectType_ACCESS_POINT
        case 0x022:
            return BACnetObjectType_ACCESS_RIGHTS
        case 0x023:
            return BACnetObjectType_ACCESS_USER
        case 0x024:
            return BACnetObjectType_ACCESS_ZONE
        case 0x025:
            return BACnetObjectType_CREDENTIAL_DATA_INPUT
        case 0x026:
            return BACnetObjectType_NETWORK_SECURITY
        case 0x027:
            return BACnetObjectType_BITSTRING_VALUE
        case 0x028:
            return BACnetObjectType_CHARACTERSTRING_VALUE
        case 0x029:
            return BACnetObjectType_DATEPATTERN_VALUE
        case 0x02A:
            return BACnetObjectType_DATE_VALUE
        case 0x02B:
            return BACnetObjectType_DATETIMEPATTERN_VALUE
        case 0x02C:
            return BACnetObjectType_DATETIME_VALUE
        case 0x02D:
            return BACnetObjectType_INTEGER_VALUE
        case 0x02E:
            return BACnetObjectType_LARGE_ANALOG_VALUE
        case 0x02F:
            return BACnetObjectType_OCTETSTRING_VALUE
        case 0x030:
            return BACnetObjectType_POSITIVE_INTEGER_VALUE
        case 0x031:
            return BACnetObjectType_TIMEPATTERN_VALUE
        case 0x032:
            return BACnetObjectType_TIME_VALUE
        case 0x033:
            return BACnetObjectType_NOTIFICATION_FORWARDER
        case 0x034:
            return BACnetObjectType_ALERT_ENROLLMENT
        case 0x035:
            return BACnetObjectType_CHANNEL
        case 0x036:
            return BACnetObjectType_LIGHTING_OUTPUT
        case 0x037:
            return BACnetObjectType_BINARY_LIGHTING_OUTPUT
        case 0x038:
            return BACnetObjectType_NETWORK_PORT
        case 0x039:
            return BACnetObjectType_ELEVATOR_GROUP
        case 0x03A:
            return BACnetObjectType_ESCALATOR
    }
    return 0
}

func CastBACnetObjectType(structType interface{}) BACnetObjectType {
    castFunc := func(typ interface{}) BACnetObjectType {
        if sBACnetObjectType, ok := typ.(BACnetObjectType); ok {
            return sBACnetObjectType
        }
        return 0
    }
    return castFunc(structType)
}

func (m BACnetObjectType) LengthInBits() uint16 {
    return 10
}

func (m BACnetObjectType) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetObjectTypeParse(io *utils.ReadBuffer) (BACnetObjectType, error) {
    val, err := io.ReadUint16(10)
    if err != nil {
        return 0, nil
    }
    return BACnetObjectTypeValueOf(val), nil
}

func (e BACnetObjectType) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint16(10, uint16(e))
    return err
}

func (e BACnetObjectType) String() string {
    switch e {
    case BACnetObjectType_ANALOG_INPUT:
        return "ANALOG_INPUT"
    case BACnetObjectType_ANALOG_OUTPUT:
        return "ANALOG_OUTPUT"
    case BACnetObjectType_ANALOG_VALUE:
        return "ANALOG_VALUE"
    case BACnetObjectType_BINARY_INPUT:
        return "BINARY_INPUT"
    case BACnetObjectType_BINARY_OUTPUT:
        return "BINARY_OUTPUT"
    case BACnetObjectType_BINARY_VALUE:
        return "BINARY_VALUE"
    case BACnetObjectType_CALENDAR:
        return "CALENDAR"
    case BACnetObjectType_COMMAND:
        return "COMMAND"
    case BACnetObjectType_DEVICE:
        return "DEVICE"
    case BACnetObjectType_EVENT_ENROLLMENT:
        return "EVENT_ENROLLMENT"
    case BACnetObjectType_FILE:
        return "FILE"
    case BACnetObjectType_GROUP:
        return "GROUP"
    case BACnetObjectType_LOOP:
        return "LOOP"
    case BACnetObjectType_MULTISTATE_INPUT:
        return "MULTISTATE_INPUT"
    case BACnetObjectType_MULTISTATE_OUTPUT:
        return "MULTISTATE_OUTPUT"
    case BACnetObjectType_NOTIFICATION_CLASS:
        return "NOTIFICATION_CLASS"
    case BACnetObjectType_PROGRAM:
        return "PROGRAM"
    case BACnetObjectType_SCHEDULE:
        return "SCHEDULE"
    case BACnetObjectType_AVERAGING:
        return "AVERAGING"
    case BACnetObjectType_MULTISTATE_VALUE:
        return "MULTISTATE_VALUE"
    case BACnetObjectType_TREND_LOG:
        return "TREND_LOG"
    case BACnetObjectType_LIFE_SAFETY_POINT:
        return "LIFE_SAFETY_POINT"
    case BACnetObjectType_LIFE_SAFETY_ZONE:
        return "LIFE_SAFETY_ZONE"
    case BACnetObjectType_ACCUMULATOR:
        return "ACCUMULATOR"
    case BACnetObjectType_PULSE_CONVERTER:
        return "PULSE_CONVERTER"
    case BACnetObjectType_EVENT_LOG:
        return "EVENT_LOG"
    case BACnetObjectType_GLOBAL_GROUP:
        return "GLOBAL_GROUP"
    case BACnetObjectType_TREND_LOG_MULTIPLE:
        return "TREND_LOG_MULTIPLE"
    case BACnetObjectType_LOAD_CONTROL:
        return "LOAD_CONTROL"
    case BACnetObjectType_STRUCTURED_VIEW:
        return "STRUCTURED_VIEW"
    case BACnetObjectType_ACCESS_DOOR:
        return "ACCESS_DOOR"
    case BACnetObjectType_TIMER:
        return "TIMER"
    case BACnetObjectType_ACCESS_CREDENTIAL:
        return "ACCESS_CREDENTIAL"
    case BACnetObjectType_ACCESS_POINT:
        return "ACCESS_POINT"
    case BACnetObjectType_ACCESS_RIGHTS:
        return "ACCESS_RIGHTS"
    case BACnetObjectType_ACCESS_USER:
        return "ACCESS_USER"
    case BACnetObjectType_ACCESS_ZONE:
        return "ACCESS_ZONE"
    case BACnetObjectType_CREDENTIAL_DATA_INPUT:
        return "CREDENTIAL_DATA_INPUT"
    case BACnetObjectType_NETWORK_SECURITY:
        return "NETWORK_SECURITY"
    case BACnetObjectType_BITSTRING_VALUE:
        return "BITSTRING_VALUE"
    case BACnetObjectType_CHARACTERSTRING_VALUE:
        return "CHARACTERSTRING_VALUE"
    case BACnetObjectType_DATEPATTERN_VALUE:
        return "DATEPATTERN_VALUE"
    case BACnetObjectType_DATE_VALUE:
        return "DATE_VALUE"
    case BACnetObjectType_DATETIMEPATTERN_VALUE:
        return "DATETIMEPATTERN_VALUE"
    case BACnetObjectType_DATETIME_VALUE:
        return "DATETIME_VALUE"
    case BACnetObjectType_INTEGER_VALUE:
        return "INTEGER_VALUE"
    case BACnetObjectType_LARGE_ANALOG_VALUE:
        return "LARGE_ANALOG_VALUE"
    case BACnetObjectType_OCTETSTRING_VALUE:
        return "OCTETSTRING_VALUE"
    case BACnetObjectType_POSITIVE_INTEGER_VALUE:
        return "POSITIVE_INTEGER_VALUE"
    case BACnetObjectType_TIMEPATTERN_VALUE:
        return "TIMEPATTERN_VALUE"
    case BACnetObjectType_TIME_VALUE:
        return "TIME_VALUE"
    case BACnetObjectType_NOTIFICATION_FORWARDER:
        return "NOTIFICATION_FORWARDER"
    case BACnetObjectType_ALERT_ENROLLMENT:
        return "ALERT_ENROLLMENT"
    case BACnetObjectType_CHANNEL:
        return "CHANNEL"
    case BACnetObjectType_LIGHTING_OUTPUT:
        return "LIGHTING_OUTPUT"
    case BACnetObjectType_BINARY_LIGHTING_OUTPUT:
        return "BINARY_LIGHTING_OUTPUT"
    case BACnetObjectType_NETWORK_PORT:
        return "NETWORK_PORT"
    case BACnetObjectType_ELEVATOR_GROUP:
        return "ELEVATOR_GROUP"
    case BACnetObjectType_ESCALATOR:
        return "ESCALATOR"
    }
    return ""
}
