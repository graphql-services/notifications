package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

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

extend type Mutation {
  seenNotification(id: ID!): Notification
  seenNotifications(principal: String!, channel: String, reference: String, referenceID: String): Boolean!
}

type Notification {
  id: ID!
  message: String!
  seen: Boolean!
  channel: String
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
  message: String!
  seen: Boolean!
  channel: String
  principal: String
  reference: String
  referenceID: String
  date: Time!
}

input NotificationUpdateInput {
  message: String
  seen: Boolean
  channel: String
  principal: String
  reference: String
  referenceID: String
  date: Time
}

enum NotificationSortType {
  ID_ASC
  ID_DESC
  MESSAGE_ASC
  MESSAGE_DESC
  SEEN_ASC
  SEEN_DESC
  CHANNEL_ASC
  CHANNEL_DESC
  PRINCIPAL_ASC
  PRINCIPAL_DESC
  REFERENCE_ASC
  REFERENCE_DESC
  REFERENCE_ID_ASC
  REFERENCE_ID_DESC
  DATE_ASC
  DATE_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
}

input NotificationFilterType {
  AND: [NotificationFilterType!]
  OR: [NotificationFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  message: String
  message_ne: String
  message_gt: String
  message_lt: String
  message_gte: String
  message_lte: String
  message_in: [String!]
  message_like: String
  message_prefix: String
  message_suffix: String
  seen: Boolean
  seen_ne: Boolean
  seen_gt: Boolean
  seen_lt: Boolean
  seen_gte: Boolean
  seen_lte: Boolean
  seen_in: [Boolean!]
  channel: String
  channel_ne: String
  channel_gt: String
  channel_lt: String
  channel_gte: String
  channel_lte: String
  channel_in: [String!]
  channel_like: String
  channel_prefix: String
  channel_suffix: String
  principal: String
  principal_ne: String
  principal_gt: String
  principal_lt: String
  principal_gte: String
  principal_lte: String
  principal_in: [String!]
  principal_like: String
  principal_prefix: String
  principal_suffix: String
  reference: String
  reference_ne: String
  reference_gt: String
  reference_lt: String
  reference_gte: String
  reference_lte: String
  reference_in: [String!]
  reference_like: String
  reference_prefix: String
  reference_suffix: String
  referenceID: String
  referenceID_ne: String
  referenceID_gt: String
  referenceID_lt: String
  referenceID_gte: String
  referenceID_lte: String
  referenceID_in: [String!]
  referenceID_like: String
  referenceID_prefix: String
  referenceID_suffix: String
  date: Time
  date_ne: Time
  date_gt: Time
  date_lt: Time
  date_gte: Time
  date_lte: Time
  date_in: [Time!]
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
}

type NotificationResultType {
  items: [Notification!]!
  count: Int!
}`
)
