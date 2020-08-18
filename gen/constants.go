package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  notification(id: ID, q: String, filter: NotificationFilterType): Notification
  notifications(offset: Int, limit: Int = 30, q: String, sort: [NotificationSortType!], filter: NotificationFilterType): NotificationResultType
}

type Mutation {
  createNotification(input: NotificationCreateInput!): Notification!
  updateNotification(id: ID!, input: NotificationUpdateInput!): Notification!
  deleteNotification(id: ID!): Notification!
  deleteAllNotifications: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

extend type Mutation {
  seenNotification(id: ID!): Notification
  seenNotifications(principal: String!, channel: String, reference: String, referenceID: String): Boolean!
}

type Notification {
  id: ID!
  groupID: ID
  message: String!
  seen: Boolean!
  channel: String
  recipient: String
  principal: String
  reference: String
  referenceID: String
  date: Time!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input NotificationCreateInput {
  id: ID
  groupID: ID
  message: String!
  seen: Boolean!
  channel: String
  recipient: String
  principal: String
  reference: String
  referenceID: String
  date: Time!
}

input NotificationUpdateInput {
  groupID: ID
  message: String
  seen: Boolean
  channel: String
  recipient: String
  principal: String
  reference: String
  referenceID: String
  date: Time
}

input NotificationSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  groupID: ObjectSortType
  groupIDMin: ObjectSortType
  groupIDMax: ObjectSortType
  message: ObjectSortType
  messageMin: ObjectSortType
  messageMax: ObjectSortType
  seen: ObjectSortType
  seenMin: ObjectSortType
  seenMax: ObjectSortType
  channel: ObjectSortType
  channelMin: ObjectSortType
  channelMax: ObjectSortType
  recipient: ObjectSortType
  recipientMin: ObjectSortType
  recipientMax: ObjectSortType
  principal: ObjectSortType
  principalMin: ObjectSortType
  principalMax: ObjectSortType
  reference: ObjectSortType
  referenceMin: ObjectSortType
  referenceMax: ObjectSortType
  referenceID: ObjectSortType
  referenceIDMin: ObjectSortType
  referenceIDMax: ObjectSortType
  date: ObjectSortType
  dateMin: ObjectSortType
  dateMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
}

input NotificationFilterType {
  AND: [NotificationFilterType!]
  OR: [NotificationFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  groupID: ID
  groupIDMin: ID
  groupIDMax: ID
  groupID_ne: ID
  groupIDMin_ne: ID
  groupIDMax_ne: ID
  groupID_gt: ID
  groupIDMin_gt: ID
  groupIDMax_gt: ID
  groupID_lt: ID
  groupIDMin_lt: ID
  groupIDMax_lt: ID
  groupID_gte: ID
  groupIDMin_gte: ID
  groupIDMax_gte: ID
  groupID_lte: ID
  groupIDMin_lte: ID
  groupIDMax_lte: ID
  groupID_in: [ID!]
  groupIDMin_in: [ID!]
  groupIDMax_in: [ID!]
  groupID_null: Boolean
  message: String
  messageMin: String
  messageMax: String
  message_ne: String
  messageMin_ne: String
  messageMax_ne: String
  message_gt: String
  messageMin_gt: String
  messageMax_gt: String
  message_lt: String
  messageMin_lt: String
  messageMax_lt: String
  message_gte: String
  messageMin_gte: String
  messageMax_gte: String
  message_lte: String
  messageMin_lte: String
  messageMax_lte: String
  message_in: [String!]
  messageMin_in: [String!]
  messageMax_in: [String!]
  message_like: String
  messageMin_like: String
  messageMax_like: String
  message_prefix: String
  messageMin_prefix: String
  messageMax_prefix: String
  message_suffix: String
  messageMin_suffix: String
  messageMax_suffix: String
  message_null: Boolean
  seen: Boolean
  seenMin: Boolean
  seenMax: Boolean
  seen_ne: Boolean
  seenMin_ne: Boolean
  seenMax_ne: Boolean
  seen_gt: Boolean
  seenMin_gt: Boolean
  seenMax_gt: Boolean
  seen_lt: Boolean
  seenMin_lt: Boolean
  seenMax_lt: Boolean
  seen_gte: Boolean
  seenMin_gte: Boolean
  seenMax_gte: Boolean
  seen_lte: Boolean
  seenMin_lte: Boolean
  seenMax_lte: Boolean
  seen_in: [Boolean!]
  seenMin_in: [Boolean!]
  seenMax_in: [Boolean!]
  seen_null: Boolean
  channel: String
  channelMin: String
  channelMax: String
  channel_ne: String
  channelMin_ne: String
  channelMax_ne: String
  channel_gt: String
  channelMin_gt: String
  channelMax_gt: String
  channel_lt: String
  channelMin_lt: String
  channelMax_lt: String
  channel_gte: String
  channelMin_gte: String
  channelMax_gte: String
  channel_lte: String
  channelMin_lte: String
  channelMax_lte: String
  channel_in: [String!]
  channelMin_in: [String!]
  channelMax_in: [String!]
  channel_like: String
  channelMin_like: String
  channelMax_like: String
  channel_prefix: String
  channelMin_prefix: String
  channelMax_prefix: String
  channel_suffix: String
  channelMin_suffix: String
  channelMax_suffix: String
  channel_null: Boolean
  recipient: String
  recipientMin: String
  recipientMax: String
  recipient_ne: String
  recipientMin_ne: String
  recipientMax_ne: String
  recipient_gt: String
  recipientMin_gt: String
  recipientMax_gt: String
  recipient_lt: String
  recipientMin_lt: String
  recipientMax_lt: String
  recipient_gte: String
  recipientMin_gte: String
  recipientMax_gte: String
  recipient_lte: String
  recipientMin_lte: String
  recipientMax_lte: String
  recipient_in: [String!]
  recipientMin_in: [String!]
  recipientMax_in: [String!]
  recipient_like: String
  recipientMin_like: String
  recipientMax_like: String
  recipient_prefix: String
  recipientMin_prefix: String
  recipientMax_prefix: String
  recipient_suffix: String
  recipientMin_suffix: String
  recipientMax_suffix: String
  recipient_null: Boolean
  principal: String
  principalMin: String
  principalMax: String
  principal_ne: String
  principalMin_ne: String
  principalMax_ne: String
  principal_gt: String
  principalMin_gt: String
  principalMax_gt: String
  principal_lt: String
  principalMin_lt: String
  principalMax_lt: String
  principal_gte: String
  principalMin_gte: String
  principalMax_gte: String
  principal_lte: String
  principalMin_lte: String
  principalMax_lte: String
  principal_in: [String!]
  principalMin_in: [String!]
  principalMax_in: [String!]
  principal_like: String
  principalMin_like: String
  principalMax_like: String
  principal_prefix: String
  principalMin_prefix: String
  principalMax_prefix: String
  principal_suffix: String
  principalMin_suffix: String
  principalMax_suffix: String
  principal_null: Boolean
  reference: String
  referenceMin: String
  referenceMax: String
  reference_ne: String
  referenceMin_ne: String
  referenceMax_ne: String
  reference_gt: String
  referenceMin_gt: String
  referenceMax_gt: String
  reference_lt: String
  referenceMin_lt: String
  referenceMax_lt: String
  reference_gte: String
  referenceMin_gte: String
  referenceMax_gte: String
  reference_lte: String
  referenceMin_lte: String
  referenceMax_lte: String
  reference_in: [String!]
  referenceMin_in: [String!]
  referenceMax_in: [String!]
  reference_like: String
  referenceMin_like: String
  referenceMax_like: String
  reference_prefix: String
  referenceMin_prefix: String
  referenceMax_prefix: String
  reference_suffix: String
  referenceMin_suffix: String
  referenceMax_suffix: String
  reference_null: Boolean
  referenceID: String
  referenceIDMin: String
  referenceIDMax: String
  referenceID_ne: String
  referenceIDMin_ne: String
  referenceIDMax_ne: String
  referenceID_gt: String
  referenceIDMin_gt: String
  referenceIDMax_gt: String
  referenceID_lt: String
  referenceIDMin_lt: String
  referenceIDMax_lt: String
  referenceID_gte: String
  referenceIDMin_gte: String
  referenceIDMax_gte: String
  referenceID_lte: String
  referenceIDMin_lte: String
  referenceIDMax_lte: String
  referenceID_in: [String!]
  referenceIDMin_in: [String!]
  referenceIDMax_in: [String!]
  referenceID_like: String
  referenceIDMin_like: String
  referenceIDMax_like: String
  referenceID_prefix: String
  referenceIDMin_prefix: String
  referenceIDMax_prefix: String
  referenceID_suffix: String
  referenceIDMin_suffix: String
  referenceIDMax_suffix: String
  referenceID_null: Boolean
  date: Time
  dateMin: Time
  dateMax: Time
  date_ne: Time
  dateMin_ne: Time
  dateMax_ne: Time
  date_gt: Time
  dateMin_gt: Time
  dateMax_gt: Time
  date_lt: Time
  dateMin_lt: Time
  dateMax_lt: Time
  date_gte: Time
  dateMin_gte: Time
  dateMax_gte: Time
  date_lte: Time
  dateMin_lte: Time
  dateMax_lte: Time
  date_in: [Time!]
  dateMin_in: [Time!]
  dateMax_in: [Time!]
  date_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
}

type NotificationResultType {
  items: [Notification!]!
  count: Int!
}`
)
