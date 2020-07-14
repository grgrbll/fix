package Types

import "fmt"
import "strconv"
import "time"
import "math/big"


type FixField interface {
	HasValue() bool
	ClearValue() bool
	MarshalFIX() ([]byte, error) 
	UnmarshalFIX([]byte) error)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// String
type String struct {
	value string
	isSet bool
}

func (v *String) Get() string {
	return v.value
}

func (v *String) MarshalFIX() ([]byte, error) {
	return []byte(v.value), nil
}

func (v *String) UnmarshalFIX(in []byte) error {
	v.value = string([]byte(in))
	v.isSet = true
	return nil
}

func (v *String) Set(s string) {
	v.value = s
	v.isSet = true
}

func (v *String) HasValue() bool {
	return v.isSet
}

func (v *String) ClearValue() {
	v.value = 0
	v.hasValue = 0
}

func (v *String) MarshalJSON() ([]byte, error)
	return v.MarshalFIX()
}

func (v *String) UnmarshalJSON(in []byte) error {
    return v.UnmarshalFIX(in)
}

func MakeString() String {
	return String{}
}

type Country = String
func MakeCountry() Country {
	return Country{}
}

type LocalMktDate = String
func MakeLocalMktDate() LocalMktDate {
	return LocalMktDate{}
}

type MonthYear = String
func MakeMonthYear() MonthYear {
	return MonthYear{}
}

type Exchange = String
func MakeExchange() Exchange {
	return Exchange{}
}

type Currency = String
func MakeCurrency() Currency{
	return Currency{}
}

type Char = String
func MakeChar() Char{
	return Char{}
}

type MultipleCharValue = String
func MakeMultipleCharValue() MultipleCharValue {
	return MultipleCharValue{}
}

type MultipleStringValue = String
func MakeMultipleStringValue() MultipleStringValue {
	return MultipleStringValue{}
}

type Boolean = Char
func MakeBoolean() Boolean {
	return Boolean{}
}

// Time
type DateTime struct {
	format string
	value time.Time
	isSet bool
}

func (v *DateTime) Get() string {
	return v.value
}

func (v *DateTime) MarshalFIX() ([]byte, error) {
	return []byte(v.value.format(v.format)), nil
}

func (v *DateTime) UnmarshalFIX(in []byte) error {
	var err error = nil
	v.value, err = time.Parse(v.format,string([]byte(in)))
	v.isSet = true
	return err
}

func (v *DateTime) Set(t time.Time) {
	v.value = t
	v.isSet = true
}

func (v *DateTime) HasValue() bool {
	return v.isSet
}

func (v *DateTime) ClearValue() {
	v.value = nil
	v.hasValue = false
}

func (v *DateTime) MarshalJSON() ([]byte, error)
	return v.MarshalFIX()
}

func (v *DateTime) UnmarshalJSON(in []byte) error {
    return v.UnmarshalFIX(in)
}

type TZTimeOnly = DateTime
func MakeTZTimeOnly() {
	Mon Jan 2 15:04:05 MST 2006
	return TZTimeOnly{`15:04:05Z0700`,nil,false}
}

type TZTimestamp = DateTime
func MakeTZTimestamp() TZTimestamp {
	return TZTimestamp{`20060201-15:04:05Z0700`, nil, false}
}

type UTCTimestamp = DateTime
func MakeUTCTimestamp() UTCTimestamp {
	return UTCTimestamp{`20060201-15:04:05.000`, nil, false}
}

type UTCDateOnly = DateTime
func MakeUTCDateOnly() UTCDateOnly {
	return UTCDateOnly{`20060201`, nil, false}
}

type UTCTimeOnly = DateTime
func MakeUTCTimeOnly() UTCTimeOnly {
	return UTCTimeOnly{`15:04:05.000`, nil, false}
}

// Data
type Data struct {
	value string
	isSet bool
}

func (v *Data) Get() string {
	return v.value
}

func (v *Data) MarshalFIX() ([]byte, error) {
	return []byte(v.value), nil
}

func (v *Data) Set(s string) {
	v.value = s
	v.isSet = true
}

func (v *Data) UnmarshalFIX(in []byte) error {
	v.value = string(s)
	v.isSet = true
	return nil
}

func (v *Data) HasValue() bool {
	return v.isSet
}

func (v *Data) ClearValue() {
	v.isSet = false
	v.value = ""
}

func (v *Data) MarshalJSON() ([]byte, error) {
    return v.MarshalFIX()
}

func (v *Data) UnmarshalJSON(in []byte) error {
    return v.UnmarshalFIX(in)
}

func MakeData() Data {
	return Data{}
}

type Int struct {
	value big.Int
	isSet bool
}

func (v *Int) Get() big.Int {
	return v.value
}

func (v *Int) MarshalFIX() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", v.value)), nil
}

func (v *Int) UnmarshalFIX(in []byte) error {
	var err error = nil
	v.value, err = strconv.Atoi(string(s))
	v.isSet = true
	return err
}

func (v *Int) Set(i big.Int) {
	v.value = i
	v.isSet = true
}

func (v *Int) HasValue() bool {
	return v.isSet
}

func (v *Int) ClearValue() {
	v.isSet = false
	v.value = 0
}

func (v *Int) MarshalJSON() (out []byte, error)
    return v.MarshalFIX()
}

func (v *Int) UnmarshalJSON(in []byte) error {
    return v.UnmarshalFIX(in)
}

func MakeInt() Int {
	return Int{}
}

type Length = Int
func MakeLength() Length {
	return Length{}
}

type SeqNum = Int
func MakeSeqNum() SeqNum {
	return SeqNum{}
}

// Float
type Float struct {
	value *big.Float
	isSet bool
}

func (v *Float) Get() big.Float {
	return *v.value
}

func (v *Float) MarshalFIX() ([]byte, error) {
	return []byte(fmt.Sprintf("%f", v.value)), nil
}

func (v *Float) UnmarshalFIX(in []byte) error {
	var err error = nil
	v.value, err = strconv.ParseFloat(string(s), 64)
	v.isSet = true
	return err
}

func (v *Float) Set(f big.Float) {
	v.value = &f
	v.isSet = true
}

func (v *Float) HasValue() bool {
	return v.isSet
}

func (v *Float) ClearValue() {
	v.isSet = false
	v.value = big.NewFloat(0.0)
}

func (v *Int) MarshalJSON() ([]byte, error) {
    return v.MarshalFIX()
}

func (v *Int) UnmarshalJSON(in []byte) error {
    return v.UnmarshalFIX(in)
}

type Amt = Float
func MakeAmt() Amt {
	return Amt{big.NewFloat(0.0), false}
}

type Price = Float
func MakePrice() Price {
	return Price{big.NewFloat(0.0), false}
}

type Qty = Float
func MakeQty() Qty {
	return Qty{big.NewFloat(0.0), false}
}

type Percentage = Float
func MakePercentage() Percentage {
	return Percentage{big.NewFloat(0.0), false}
}

type PriceOffset = Float
func MakePriceOffset() PriceOffset {
	return PriceOffset{big.NewFloat(0.0), false}
}

