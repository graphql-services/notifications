// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type NotificationFilterType struct {
	And                  []*NotificationFilterType `json:"AND"`
	Or                   []*NotificationFilterType `json:"OR"`
	ID                   *string                   `json:"id"`
	IDMin                *string                   `json:"idMin"`
	IDMax                *string                   `json:"idMax"`
	IDNe                 *string                   `json:"id_ne"`
	IDMinNe              *string                   `json:"idMin_ne"`
	IDMaxNe              *string                   `json:"idMax_ne"`
	IDGt                 *string                   `json:"id_gt"`
	IDMinGt              *string                   `json:"idMin_gt"`
	IDMaxGt              *string                   `json:"idMax_gt"`
	IDLt                 *string                   `json:"id_lt"`
	IDMinLt              *string                   `json:"idMin_lt"`
	IDMaxLt              *string                   `json:"idMax_lt"`
	IDGte                *string                   `json:"id_gte"`
	IDMinGte             *string                   `json:"idMin_gte"`
	IDMaxGte             *string                   `json:"idMax_gte"`
	IDLte                *string                   `json:"id_lte"`
	IDMinLte             *string                   `json:"idMin_lte"`
	IDMaxLte             *string                   `json:"idMax_lte"`
	IDIn                 []string                  `json:"id_in"`
	IDMinIn              []string                  `json:"idMin_in"`
	IDMaxIn              []string                  `json:"idMax_in"`
	IDNull               *bool                     `json:"id_null"`
	GroupID              *string                   `json:"groupID"`
	GroupIDMin           *string                   `json:"groupIDMin"`
	GroupIDMax           *string                   `json:"groupIDMax"`
	GroupIDNe            *string                   `json:"groupID_ne"`
	GroupIDMinNe         *string                   `json:"groupIDMin_ne"`
	GroupIDMaxNe         *string                   `json:"groupIDMax_ne"`
	GroupIDGt            *string                   `json:"groupID_gt"`
	GroupIDMinGt         *string                   `json:"groupIDMin_gt"`
	GroupIDMaxGt         *string                   `json:"groupIDMax_gt"`
	GroupIDLt            *string                   `json:"groupID_lt"`
	GroupIDMinLt         *string                   `json:"groupIDMin_lt"`
	GroupIDMaxLt         *string                   `json:"groupIDMax_lt"`
	GroupIDGte           *string                   `json:"groupID_gte"`
	GroupIDMinGte        *string                   `json:"groupIDMin_gte"`
	GroupIDMaxGte        *string                   `json:"groupIDMax_gte"`
	GroupIDLte           *string                   `json:"groupID_lte"`
	GroupIDMinLte        *string                   `json:"groupIDMin_lte"`
	GroupIDMaxLte        *string                   `json:"groupIDMax_lte"`
	GroupIDIn            []string                  `json:"groupID_in"`
	GroupIDMinIn         []string                  `json:"groupIDMin_in"`
	GroupIDMaxIn         []string                  `json:"groupIDMax_in"`
	GroupIDNull          *bool                     `json:"groupID_null"`
	Subject              *string                   `json:"subject"`
	SubjectMin           *string                   `json:"subjectMin"`
	SubjectMax           *string                   `json:"subjectMax"`
	SubjectNe            *string                   `json:"subject_ne"`
	SubjectMinNe         *string                   `json:"subjectMin_ne"`
	SubjectMaxNe         *string                   `json:"subjectMax_ne"`
	SubjectGt            *string                   `json:"subject_gt"`
	SubjectMinGt         *string                   `json:"subjectMin_gt"`
	SubjectMaxGt         *string                   `json:"subjectMax_gt"`
	SubjectLt            *string                   `json:"subject_lt"`
	SubjectMinLt         *string                   `json:"subjectMin_lt"`
	SubjectMaxLt         *string                   `json:"subjectMax_lt"`
	SubjectGte           *string                   `json:"subject_gte"`
	SubjectMinGte        *string                   `json:"subjectMin_gte"`
	SubjectMaxGte        *string                   `json:"subjectMax_gte"`
	SubjectLte           *string                   `json:"subject_lte"`
	SubjectMinLte        *string                   `json:"subjectMin_lte"`
	SubjectMaxLte        *string                   `json:"subjectMax_lte"`
	SubjectIn            []string                  `json:"subject_in"`
	SubjectMinIn         []string                  `json:"subjectMin_in"`
	SubjectMaxIn         []string                  `json:"subjectMax_in"`
	SubjectLike          *string                   `json:"subject_like"`
	SubjectMinLike       *string                   `json:"subjectMin_like"`
	SubjectMaxLike       *string                   `json:"subjectMax_like"`
	SubjectPrefix        *string                   `json:"subject_prefix"`
	SubjectMinPrefix     *string                   `json:"subjectMin_prefix"`
	SubjectMaxPrefix     *string                   `json:"subjectMax_prefix"`
	SubjectSuffix        *string                   `json:"subject_suffix"`
	SubjectMinSuffix     *string                   `json:"subjectMin_suffix"`
	SubjectMaxSuffix     *string                   `json:"subjectMax_suffix"`
	SubjectNull          *bool                     `json:"subject_null"`
	Message              *string                   `json:"message"`
	MessageMin           *string                   `json:"messageMin"`
	MessageMax           *string                   `json:"messageMax"`
	MessageNe            *string                   `json:"message_ne"`
	MessageMinNe         *string                   `json:"messageMin_ne"`
	MessageMaxNe         *string                   `json:"messageMax_ne"`
	MessageGt            *string                   `json:"message_gt"`
	MessageMinGt         *string                   `json:"messageMin_gt"`
	MessageMaxGt         *string                   `json:"messageMax_gt"`
	MessageLt            *string                   `json:"message_lt"`
	MessageMinLt         *string                   `json:"messageMin_lt"`
	MessageMaxLt         *string                   `json:"messageMax_lt"`
	MessageGte           *string                   `json:"message_gte"`
	MessageMinGte        *string                   `json:"messageMin_gte"`
	MessageMaxGte        *string                   `json:"messageMax_gte"`
	MessageLte           *string                   `json:"message_lte"`
	MessageMinLte        *string                   `json:"messageMin_lte"`
	MessageMaxLte        *string                   `json:"messageMax_lte"`
	MessageIn            []string                  `json:"message_in"`
	MessageMinIn         []string                  `json:"messageMin_in"`
	MessageMaxIn         []string                  `json:"messageMax_in"`
	MessageLike          *string                   `json:"message_like"`
	MessageMinLike       *string                   `json:"messageMin_like"`
	MessageMaxLike       *string                   `json:"messageMax_like"`
	MessagePrefix        *string                   `json:"message_prefix"`
	MessageMinPrefix     *string                   `json:"messageMin_prefix"`
	MessageMaxPrefix     *string                   `json:"messageMax_prefix"`
	MessageSuffix        *string                   `json:"message_suffix"`
	MessageMinSuffix     *string                   `json:"messageMin_suffix"`
	MessageMaxSuffix     *string                   `json:"messageMax_suffix"`
	MessageNull          *bool                     `json:"message_null"`
	Seen                 *bool                     `json:"seen"`
	SeenMin              *bool                     `json:"seenMin"`
	SeenMax              *bool                     `json:"seenMax"`
	SeenNe               *bool                     `json:"seen_ne"`
	SeenMinNe            *bool                     `json:"seenMin_ne"`
	SeenMaxNe            *bool                     `json:"seenMax_ne"`
	SeenGt               *bool                     `json:"seen_gt"`
	SeenMinGt            *bool                     `json:"seenMin_gt"`
	SeenMaxGt            *bool                     `json:"seenMax_gt"`
	SeenLt               *bool                     `json:"seen_lt"`
	SeenMinLt            *bool                     `json:"seenMin_lt"`
	SeenMaxLt            *bool                     `json:"seenMax_lt"`
	SeenGte              *bool                     `json:"seen_gte"`
	SeenMinGte           *bool                     `json:"seenMin_gte"`
	SeenMaxGte           *bool                     `json:"seenMax_gte"`
	SeenLte              *bool                     `json:"seen_lte"`
	SeenMinLte           *bool                     `json:"seenMin_lte"`
	SeenMaxLte           *bool                     `json:"seenMax_lte"`
	SeenIn               []bool                    `json:"seen_in"`
	SeenMinIn            []bool                    `json:"seenMin_in"`
	SeenMaxIn            []bool                    `json:"seenMax_in"`
	SeenNull             *bool                     `json:"seen_null"`
	Channel              *string                   `json:"channel"`
	ChannelMin           *string                   `json:"channelMin"`
	ChannelMax           *string                   `json:"channelMax"`
	ChannelNe            *string                   `json:"channel_ne"`
	ChannelMinNe         *string                   `json:"channelMin_ne"`
	ChannelMaxNe         *string                   `json:"channelMax_ne"`
	ChannelGt            *string                   `json:"channel_gt"`
	ChannelMinGt         *string                   `json:"channelMin_gt"`
	ChannelMaxGt         *string                   `json:"channelMax_gt"`
	ChannelLt            *string                   `json:"channel_lt"`
	ChannelMinLt         *string                   `json:"channelMin_lt"`
	ChannelMaxLt         *string                   `json:"channelMax_lt"`
	ChannelGte           *string                   `json:"channel_gte"`
	ChannelMinGte        *string                   `json:"channelMin_gte"`
	ChannelMaxGte        *string                   `json:"channelMax_gte"`
	ChannelLte           *string                   `json:"channel_lte"`
	ChannelMinLte        *string                   `json:"channelMin_lte"`
	ChannelMaxLte        *string                   `json:"channelMax_lte"`
	ChannelIn            []string                  `json:"channel_in"`
	ChannelMinIn         []string                  `json:"channelMin_in"`
	ChannelMaxIn         []string                  `json:"channelMax_in"`
	ChannelLike          *string                   `json:"channel_like"`
	ChannelMinLike       *string                   `json:"channelMin_like"`
	ChannelMaxLike       *string                   `json:"channelMax_like"`
	ChannelPrefix        *string                   `json:"channel_prefix"`
	ChannelMinPrefix     *string                   `json:"channelMin_prefix"`
	ChannelMaxPrefix     *string                   `json:"channelMax_prefix"`
	ChannelSuffix        *string                   `json:"channel_suffix"`
	ChannelMinSuffix     *string                   `json:"channelMin_suffix"`
	ChannelMaxSuffix     *string                   `json:"channelMax_suffix"`
	ChannelNull          *bool                     `json:"channel_null"`
	Recipient            *string                   `json:"recipient"`
	RecipientMin         *string                   `json:"recipientMin"`
	RecipientMax         *string                   `json:"recipientMax"`
	RecipientNe          *string                   `json:"recipient_ne"`
	RecipientMinNe       *string                   `json:"recipientMin_ne"`
	RecipientMaxNe       *string                   `json:"recipientMax_ne"`
	RecipientGt          *string                   `json:"recipient_gt"`
	RecipientMinGt       *string                   `json:"recipientMin_gt"`
	RecipientMaxGt       *string                   `json:"recipientMax_gt"`
	RecipientLt          *string                   `json:"recipient_lt"`
	RecipientMinLt       *string                   `json:"recipientMin_lt"`
	RecipientMaxLt       *string                   `json:"recipientMax_lt"`
	RecipientGte         *string                   `json:"recipient_gte"`
	RecipientMinGte      *string                   `json:"recipientMin_gte"`
	RecipientMaxGte      *string                   `json:"recipientMax_gte"`
	RecipientLte         *string                   `json:"recipient_lte"`
	RecipientMinLte      *string                   `json:"recipientMin_lte"`
	RecipientMaxLte      *string                   `json:"recipientMax_lte"`
	RecipientIn          []string                  `json:"recipient_in"`
	RecipientMinIn       []string                  `json:"recipientMin_in"`
	RecipientMaxIn       []string                  `json:"recipientMax_in"`
	RecipientLike        *string                   `json:"recipient_like"`
	RecipientMinLike     *string                   `json:"recipientMin_like"`
	RecipientMaxLike     *string                   `json:"recipientMax_like"`
	RecipientPrefix      *string                   `json:"recipient_prefix"`
	RecipientMinPrefix   *string                   `json:"recipientMin_prefix"`
	RecipientMaxPrefix   *string                   `json:"recipientMax_prefix"`
	RecipientSuffix      *string                   `json:"recipient_suffix"`
	RecipientMinSuffix   *string                   `json:"recipientMin_suffix"`
	RecipientMaxSuffix   *string                   `json:"recipientMax_suffix"`
	RecipientNull        *bool                     `json:"recipient_null"`
	Principal            *string                   `json:"principal"`
	PrincipalMin         *string                   `json:"principalMin"`
	PrincipalMax         *string                   `json:"principalMax"`
	PrincipalNe          *string                   `json:"principal_ne"`
	PrincipalMinNe       *string                   `json:"principalMin_ne"`
	PrincipalMaxNe       *string                   `json:"principalMax_ne"`
	PrincipalGt          *string                   `json:"principal_gt"`
	PrincipalMinGt       *string                   `json:"principalMin_gt"`
	PrincipalMaxGt       *string                   `json:"principalMax_gt"`
	PrincipalLt          *string                   `json:"principal_lt"`
	PrincipalMinLt       *string                   `json:"principalMin_lt"`
	PrincipalMaxLt       *string                   `json:"principalMax_lt"`
	PrincipalGte         *string                   `json:"principal_gte"`
	PrincipalMinGte      *string                   `json:"principalMin_gte"`
	PrincipalMaxGte      *string                   `json:"principalMax_gte"`
	PrincipalLte         *string                   `json:"principal_lte"`
	PrincipalMinLte      *string                   `json:"principalMin_lte"`
	PrincipalMaxLte      *string                   `json:"principalMax_lte"`
	PrincipalIn          []string                  `json:"principal_in"`
	PrincipalMinIn       []string                  `json:"principalMin_in"`
	PrincipalMaxIn       []string                  `json:"principalMax_in"`
	PrincipalLike        *string                   `json:"principal_like"`
	PrincipalMinLike     *string                   `json:"principalMin_like"`
	PrincipalMaxLike     *string                   `json:"principalMax_like"`
	PrincipalPrefix      *string                   `json:"principal_prefix"`
	PrincipalMinPrefix   *string                   `json:"principalMin_prefix"`
	PrincipalMaxPrefix   *string                   `json:"principalMax_prefix"`
	PrincipalSuffix      *string                   `json:"principal_suffix"`
	PrincipalMinSuffix   *string                   `json:"principalMin_suffix"`
	PrincipalMaxSuffix   *string                   `json:"principalMax_suffix"`
	PrincipalNull        *bool                     `json:"principal_null"`
	Reference            *string                   `json:"reference"`
	ReferenceMin         *string                   `json:"referenceMin"`
	ReferenceMax         *string                   `json:"referenceMax"`
	ReferenceNe          *string                   `json:"reference_ne"`
	ReferenceMinNe       *string                   `json:"referenceMin_ne"`
	ReferenceMaxNe       *string                   `json:"referenceMax_ne"`
	ReferenceGt          *string                   `json:"reference_gt"`
	ReferenceMinGt       *string                   `json:"referenceMin_gt"`
	ReferenceMaxGt       *string                   `json:"referenceMax_gt"`
	ReferenceLt          *string                   `json:"reference_lt"`
	ReferenceMinLt       *string                   `json:"referenceMin_lt"`
	ReferenceMaxLt       *string                   `json:"referenceMax_lt"`
	ReferenceGte         *string                   `json:"reference_gte"`
	ReferenceMinGte      *string                   `json:"referenceMin_gte"`
	ReferenceMaxGte      *string                   `json:"referenceMax_gte"`
	ReferenceLte         *string                   `json:"reference_lte"`
	ReferenceMinLte      *string                   `json:"referenceMin_lte"`
	ReferenceMaxLte      *string                   `json:"referenceMax_lte"`
	ReferenceIn          []string                  `json:"reference_in"`
	ReferenceMinIn       []string                  `json:"referenceMin_in"`
	ReferenceMaxIn       []string                  `json:"referenceMax_in"`
	ReferenceLike        *string                   `json:"reference_like"`
	ReferenceMinLike     *string                   `json:"referenceMin_like"`
	ReferenceMaxLike     *string                   `json:"referenceMax_like"`
	ReferencePrefix      *string                   `json:"reference_prefix"`
	ReferenceMinPrefix   *string                   `json:"referenceMin_prefix"`
	ReferenceMaxPrefix   *string                   `json:"referenceMax_prefix"`
	ReferenceSuffix      *string                   `json:"reference_suffix"`
	ReferenceMinSuffix   *string                   `json:"referenceMin_suffix"`
	ReferenceMaxSuffix   *string                   `json:"referenceMax_suffix"`
	ReferenceNull        *bool                     `json:"reference_null"`
	ReferenceID          *string                   `json:"referenceID"`
	ReferenceIDMin       *string                   `json:"referenceIDMin"`
	ReferenceIDMax       *string                   `json:"referenceIDMax"`
	ReferenceIDNe        *string                   `json:"referenceID_ne"`
	ReferenceIDMinNe     *string                   `json:"referenceIDMin_ne"`
	ReferenceIDMaxNe     *string                   `json:"referenceIDMax_ne"`
	ReferenceIDGt        *string                   `json:"referenceID_gt"`
	ReferenceIDMinGt     *string                   `json:"referenceIDMin_gt"`
	ReferenceIDMaxGt     *string                   `json:"referenceIDMax_gt"`
	ReferenceIDLt        *string                   `json:"referenceID_lt"`
	ReferenceIDMinLt     *string                   `json:"referenceIDMin_lt"`
	ReferenceIDMaxLt     *string                   `json:"referenceIDMax_lt"`
	ReferenceIDGte       *string                   `json:"referenceID_gte"`
	ReferenceIDMinGte    *string                   `json:"referenceIDMin_gte"`
	ReferenceIDMaxGte    *string                   `json:"referenceIDMax_gte"`
	ReferenceIDLte       *string                   `json:"referenceID_lte"`
	ReferenceIDMinLte    *string                   `json:"referenceIDMin_lte"`
	ReferenceIDMaxLte    *string                   `json:"referenceIDMax_lte"`
	ReferenceIDIn        []string                  `json:"referenceID_in"`
	ReferenceIDMinIn     []string                  `json:"referenceIDMin_in"`
	ReferenceIDMaxIn     []string                  `json:"referenceIDMax_in"`
	ReferenceIDLike      *string                   `json:"referenceID_like"`
	ReferenceIDMinLike   *string                   `json:"referenceIDMin_like"`
	ReferenceIDMaxLike   *string                   `json:"referenceIDMax_like"`
	ReferenceIDPrefix    *string                   `json:"referenceID_prefix"`
	ReferenceIDMinPrefix *string                   `json:"referenceIDMin_prefix"`
	ReferenceIDMaxPrefix *string                   `json:"referenceIDMax_prefix"`
	ReferenceIDSuffix    *string                   `json:"referenceID_suffix"`
	ReferenceIDMinSuffix *string                   `json:"referenceIDMin_suffix"`
	ReferenceIDMaxSuffix *string                   `json:"referenceIDMax_suffix"`
	ReferenceIDNull      *bool                     `json:"referenceID_null"`
	Date                 *time.Time                `json:"date"`
	DateMin              *time.Time                `json:"dateMin"`
	DateMax              *time.Time                `json:"dateMax"`
	DateNe               *time.Time                `json:"date_ne"`
	DateMinNe            *time.Time                `json:"dateMin_ne"`
	DateMaxNe            *time.Time                `json:"dateMax_ne"`
	DateGt               *time.Time                `json:"date_gt"`
	DateMinGt            *time.Time                `json:"dateMin_gt"`
	DateMaxGt            *time.Time                `json:"dateMax_gt"`
	DateLt               *time.Time                `json:"date_lt"`
	DateMinLt            *time.Time                `json:"dateMin_lt"`
	DateMaxLt            *time.Time                `json:"dateMax_lt"`
	DateGte              *time.Time                `json:"date_gte"`
	DateMinGte           *time.Time                `json:"dateMin_gte"`
	DateMaxGte           *time.Time                `json:"dateMax_gte"`
	DateLte              *time.Time                `json:"date_lte"`
	DateMinLte           *time.Time                `json:"dateMin_lte"`
	DateMaxLte           *time.Time                `json:"dateMax_lte"`
	DateIn               []*time.Time              `json:"date_in"`
	DateMinIn            []*time.Time              `json:"dateMin_in"`
	DateMaxIn            []*time.Time              `json:"dateMax_in"`
	DateNull             *bool                     `json:"date_null"`
	UpdatedAt            *time.Time                `json:"updatedAt"`
	UpdatedAtMin         *time.Time                `json:"updatedAtMin"`
	UpdatedAtMax         *time.Time                `json:"updatedAtMax"`
	UpdatedAtNe          *time.Time                `json:"updatedAt_ne"`
	UpdatedAtMinNe       *time.Time                `json:"updatedAtMin_ne"`
	UpdatedAtMaxNe       *time.Time                `json:"updatedAtMax_ne"`
	UpdatedAtGt          *time.Time                `json:"updatedAt_gt"`
	UpdatedAtMinGt       *time.Time                `json:"updatedAtMin_gt"`
	UpdatedAtMaxGt       *time.Time                `json:"updatedAtMax_gt"`
	UpdatedAtLt          *time.Time                `json:"updatedAt_lt"`
	UpdatedAtMinLt       *time.Time                `json:"updatedAtMin_lt"`
	UpdatedAtMaxLt       *time.Time                `json:"updatedAtMax_lt"`
	UpdatedAtGte         *time.Time                `json:"updatedAt_gte"`
	UpdatedAtMinGte      *time.Time                `json:"updatedAtMin_gte"`
	UpdatedAtMaxGte      *time.Time                `json:"updatedAtMax_gte"`
	UpdatedAtLte         *time.Time                `json:"updatedAt_lte"`
	UpdatedAtMinLte      *time.Time                `json:"updatedAtMin_lte"`
	UpdatedAtMaxLte      *time.Time                `json:"updatedAtMax_lte"`
	UpdatedAtIn          []*time.Time              `json:"updatedAt_in"`
	UpdatedAtMinIn       []*time.Time              `json:"updatedAtMin_in"`
	UpdatedAtMaxIn       []*time.Time              `json:"updatedAtMax_in"`
	UpdatedAtNull        *bool                     `json:"updatedAt_null"`
	CreatedAt            *time.Time                `json:"createdAt"`
	CreatedAtMin         *time.Time                `json:"createdAtMin"`
	CreatedAtMax         *time.Time                `json:"createdAtMax"`
	CreatedAtNe          *time.Time                `json:"createdAt_ne"`
	CreatedAtMinNe       *time.Time                `json:"createdAtMin_ne"`
	CreatedAtMaxNe       *time.Time                `json:"createdAtMax_ne"`
	CreatedAtGt          *time.Time                `json:"createdAt_gt"`
	CreatedAtMinGt       *time.Time                `json:"createdAtMin_gt"`
	CreatedAtMaxGt       *time.Time                `json:"createdAtMax_gt"`
	CreatedAtLt          *time.Time                `json:"createdAt_lt"`
	CreatedAtMinLt       *time.Time                `json:"createdAtMin_lt"`
	CreatedAtMaxLt       *time.Time                `json:"createdAtMax_lt"`
	CreatedAtGte         *time.Time                `json:"createdAt_gte"`
	CreatedAtMinGte      *time.Time                `json:"createdAtMin_gte"`
	CreatedAtMaxGte      *time.Time                `json:"createdAtMax_gte"`
	CreatedAtLte         *time.Time                `json:"createdAt_lte"`
	CreatedAtMinLte      *time.Time                `json:"createdAtMin_lte"`
	CreatedAtMaxLte      *time.Time                `json:"createdAtMax_lte"`
	CreatedAtIn          []*time.Time              `json:"createdAt_in"`
	CreatedAtMinIn       []*time.Time              `json:"createdAtMin_in"`
	CreatedAtMaxIn       []*time.Time              `json:"createdAtMax_in"`
	CreatedAtNull        *bool                     `json:"createdAt_null"`
	UpdatedBy            *string                   `json:"updatedBy"`
	UpdatedByMin         *string                   `json:"updatedByMin"`
	UpdatedByMax         *string                   `json:"updatedByMax"`
	UpdatedByNe          *string                   `json:"updatedBy_ne"`
	UpdatedByMinNe       *string                   `json:"updatedByMin_ne"`
	UpdatedByMaxNe       *string                   `json:"updatedByMax_ne"`
	UpdatedByGt          *string                   `json:"updatedBy_gt"`
	UpdatedByMinGt       *string                   `json:"updatedByMin_gt"`
	UpdatedByMaxGt       *string                   `json:"updatedByMax_gt"`
	UpdatedByLt          *string                   `json:"updatedBy_lt"`
	UpdatedByMinLt       *string                   `json:"updatedByMin_lt"`
	UpdatedByMaxLt       *string                   `json:"updatedByMax_lt"`
	UpdatedByGte         *string                   `json:"updatedBy_gte"`
	UpdatedByMinGte      *string                   `json:"updatedByMin_gte"`
	UpdatedByMaxGte      *string                   `json:"updatedByMax_gte"`
	UpdatedByLte         *string                   `json:"updatedBy_lte"`
	UpdatedByMinLte      *string                   `json:"updatedByMin_lte"`
	UpdatedByMaxLte      *string                   `json:"updatedByMax_lte"`
	UpdatedByIn          []string                  `json:"updatedBy_in"`
	UpdatedByMinIn       []string                  `json:"updatedByMin_in"`
	UpdatedByMaxIn       []string                  `json:"updatedByMax_in"`
	UpdatedByNull        *bool                     `json:"updatedBy_null"`
	CreatedBy            *string                   `json:"createdBy"`
	CreatedByMin         *string                   `json:"createdByMin"`
	CreatedByMax         *string                   `json:"createdByMax"`
	CreatedByNe          *string                   `json:"createdBy_ne"`
	CreatedByMinNe       *string                   `json:"createdByMin_ne"`
	CreatedByMaxNe       *string                   `json:"createdByMax_ne"`
	CreatedByGt          *string                   `json:"createdBy_gt"`
	CreatedByMinGt       *string                   `json:"createdByMin_gt"`
	CreatedByMaxGt       *string                   `json:"createdByMax_gt"`
	CreatedByLt          *string                   `json:"createdBy_lt"`
	CreatedByMinLt       *string                   `json:"createdByMin_lt"`
	CreatedByMaxLt       *string                   `json:"createdByMax_lt"`
	CreatedByGte         *string                   `json:"createdBy_gte"`
	CreatedByMinGte      *string                   `json:"createdByMin_gte"`
	CreatedByMaxGte      *string                   `json:"createdByMax_gte"`
	CreatedByLte         *string                   `json:"createdBy_lte"`
	CreatedByMinLte      *string                   `json:"createdByMin_lte"`
	CreatedByMaxLte      *string                   `json:"createdByMax_lte"`
	CreatedByIn          []string                  `json:"createdBy_in"`
	CreatedByMinIn       []string                  `json:"createdByMin_in"`
	CreatedByMaxIn       []string                  `json:"createdByMax_in"`
	CreatedByNull        *bool                     `json:"createdBy_null"`
}

type NotificationSortType struct {
	ID             *ObjectSortType `json:"id"`
	IDMin          *ObjectSortType `json:"idMin"`
	IDMax          *ObjectSortType `json:"idMax"`
	GroupID        *ObjectSortType `json:"groupID"`
	GroupIDMin     *ObjectSortType `json:"groupIDMin"`
	GroupIDMax     *ObjectSortType `json:"groupIDMax"`
	Subject        *ObjectSortType `json:"subject"`
	SubjectMin     *ObjectSortType `json:"subjectMin"`
	SubjectMax     *ObjectSortType `json:"subjectMax"`
	Message        *ObjectSortType `json:"message"`
	MessageMin     *ObjectSortType `json:"messageMin"`
	MessageMax     *ObjectSortType `json:"messageMax"`
	Seen           *ObjectSortType `json:"seen"`
	SeenMin        *ObjectSortType `json:"seenMin"`
	SeenMax        *ObjectSortType `json:"seenMax"`
	Channel        *ObjectSortType `json:"channel"`
	ChannelMin     *ObjectSortType `json:"channelMin"`
	ChannelMax     *ObjectSortType `json:"channelMax"`
	Recipient      *ObjectSortType `json:"recipient"`
	RecipientMin   *ObjectSortType `json:"recipientMin"`
	RecipientMax   *ObjectSortType `json:"recipientMax"`
	Principal      *ObjectSortType `json:"principal"`
	PrincipalMin   *ObjectSortType `json:"principalMin"`
	PrincipalMax   *ObjectSortType `json:"principalMax"`
	Reference      *ObjectSortType `json:"reference"`
	ReferenceMin   *ObjectSortType `json:"referenceMin"`
	ReferenceMax   *ObjectSortType `json:"referenceMax"`
	ReferenceID    *ObjectSortType `json:"referenceID"`
	ReferenceIDMin *ObjectSortType `json:"referenceIDMin"`
	ReferenceIDMax *ObjectSortType `json:"referenceIDMax"`
	Date           *ObjectSortType `json:"date"`
	DateMin        *ObjectSortType `json:"dateMin"`
	DateMax        *ObjectSortType `json:"dateMax"`
	UpdatedAt      *ObjectSortType `json:"updatedAt"`
	UpdatedAtMin   *ObjectSortType `json:"updatedAtMin"`
	UpdatedAtMax   *ObjectSortType `json:"updatedAtMax"`
	CreatedAt      *ObjectSortType `json:"createdAt"`
	CreatedAtMin   *ObjectSortType `json:"createdAtMin"`
	CreatedAtMax   *ObjectSortType `json:"createdAtMax"`
	UpdatedBy      *ObjectSortType `json:"updatedBy"`
	UpdatedByMin   *ObjectSortType `json:"updatedByMin"`
	UpdatedByMax   *ObjectSortType `json:"updatedByMax"`
	CreatedBy      *ObjectSortType `json:"createdBy"`
	CreatedByMin   *ObjectSortType `json:"createdByMin"`
	CreatedByMax   *ObjectSortType `json:"createdByMax"`
}

type _Service struct {
	Sdl *string `json:"sdl"`
}

type ObjectSortType string

const (
	ObjectSortTypeAsc  ObjectSortType = "ASC"
	ObjectSortTypeDesc ObjectSortType = "DESC"
)

var AllObjectSortType = []ObjectSortType{
	ObjectSortTypeAsc,
	ObjectSortTypeDesc,
}

func (e ObjectSortType) IsValid() bool {
	switch e {
	case ObjectSortTypeAsc, ObjectSortTypeDesc:
		return true
	}
	return false
}

func (e ObjectSortType) String() string {
	return string(e)
}

func (e *ObjectSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectSortType", str)
	}
	return nil
}

func (e ObjectSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
