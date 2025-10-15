/* eslint-disable */

export interface AddressRow {
  address1: string;
  address2: string;
  city: string;
  country: string;
  createdAt: string;
  defaultDelivery: string;
  defaultPickup: string;
  id: number;
  isTest: string;
  lat: number;
  lon: number;
  name: string;
  phone: string;
  state: string;
  status: string;
  updatedAt: string;
  userId: number;
  zipcode: string;
}

export const AddressRowTable = "addresses" as const;

export const AddressRowColumns = {
  address1: "address1",
  address2: "address2",
  city: "city",
  country: "country",
  createdAt: "created_at",
  defaultDelivery: "default_delivery",
  defaultPickup: "default_pickup",
  id: "id",
  isTest: "is_test",
  lat: "lat",
  lon: "lon",
  name: "name",
  phone: "phone",
  state: "state",
  status: "status",
  updatedAt: "updated_at",
  userId: "user_id",
  zipcode: "zipcode",
} as const;

export type AddressRowColumn = keyof typeof AddressRowColumns;

export const AddressRowOrderBy = {
  address1: "address1",
  address2: "address2",
  city: "city",
  country: "country",
  createdAt: "created_at",
  defaultDelivery: "default_delivery",
  defaultPickup: "default_pickup",
  id: "id",
  isTest: "is_test",
  lat: "lat",
  lon: "lon",
  name: "name",
  phone: "phone",
  state: "state",
  status: "status",
  updatedAt: "updated_at",
  userId: "user_id",
  zipcode: "zipcode",
} as const;

export interface CacheRow {
  expiration: number;
  key: string;
  value: string;
}

export const CacheRowTable = "cache" as const;

export const CacheRowColumns = {
  expiration: "expiration",
  key: "key",
  value: "value",
} as const;

export type CacheRowColumn = keyof typeof CacheRowColumns;

export const CacheRowOrderBy = {
  expiration: "expiration",
  key: "key",
  value: "value",
} as const;

export interface CategoryRow {
  createdAt: string;
  id: number;
  name: string;
  position: number;
  slug: string;
  status: string;
  storeId: number;
  updatedAt: string;
}

export const CategoryRowTable = "categories" as const;

export const CategoryRowColumns = {
  createdAt: "created_at",
  id: "id",
  name: "name",
  position: "position",
  slug: "slug",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export type CategoryRowColumn = keyof typeof CategoryRowColumns;

export const CategoryRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  name: "name",
  position: "position",
  slug: "slug",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export interface CategoryProductRow {
  categoryId: number;
  createdAt: string;
  id: number;
  position: number;
  productId: number;
  storeId: number;
  updatedAt: string;
}

export const CategoryProductRowTable = "category_products" as const;

export const CategoryProductRowColumns = {
  categoryId: "category_id",
  createdAt: "created_at",
  id: "id",
  position: "position",
  productId: "product_id",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export type CategoryProductRowColumn = keyof typeof CategoryProductRowColumns;

export const CategoryProductRowOrderBy = {
  categoryId: "category_id",
  createdAt: "created_at",
  id: "id",
  position: "position",
  productId: "product_id",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export interface CourierLocationRow {
  accuracy: number;
  altitude: number;
  altitudeAccuracy: number;
  courierUuid: string;
  heading: number;
  id: number;
  lat: number;
  locationTime: string;
  lon: number;
  speed: number;
}

export const CourierLocationRowTable = "courier_locations" as const;

export const CourierLocationRowColumns = {
  accuracy: "accuracy",
  altitude: "altitude",
  altitudeAccuracy: "altitude_accuracy",
  courierUuid: "courier_uuid",
  heading: "heading",
  id: "id",
  lat: "lat",
  locationTime: "location_time",
  lon: "lon",
  speed: "speed",
} as const;

export type CourierLocationRowColumn = keyof typeof CourierLocationRowColumns;

export const CourierLocationRowOrderBy = {
  accuracy: "accuracy",
  altitude: "altitude",
  altitudeAccuracy: "altitude_accuracy",
  courierUuid: "courier_uuid",
  heading: "heading",
  id: "id",
  lat: "lat",
  locationTime: "location_time",
  lon: "lon",
  speed: "speed",
} as const;

export interface CourierOrderRow {
  courierUuid: string;
  createdAt: string;
  id: number;
  orderId: number;
  status: string;
  updatedAt: string;
}

export const CourierOrderRowTable = "courier_orders" as const;

export const CourierOrderRowColumns = {
  courierUuid: "courier_uuid",
  createdAt: "created_at",
  id: "id",
  orderId: "order_id",
  status: "status",
  updatedAt: "updated_at",
} as const;

export type CourierOrderRowColumn = keyof typeof CourierOrderRowColumns;

export const CourierOrderRowOrderBy = {
  courierUuid: "courier_uuid",
  createdAt: "created_at",
  id: "id",
  orderId: "order_id",
  status: "status",
  updatedAt: "updated_at",
} as const;

export interface CustomizationSearchRow {
  createdAt: string;
  customizationFullName: string;
  id: number;
  name: string;
  relationId: number;
  updatedAt: string;
  userId: number;
}

export const CustomizationSearchRowTable = "customization_search" as const;

export const CustomizationSearchRowColumns = {
  createdAt: "created_at",
  customizationFullName: "customization_full_name",
  id: "id",
  name: "name",
  relationId: "relation_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type CustomizationSearchRowColumn = keyof typeof CustomizationSearchRowColumns;

export const CustomizationSearchRowOrderBy = {
  createdAt: "created_at",
  customizationFullName: "customization_full_name",
  id: "id",
  name: "name",
  relationId: "relation_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface DiscountInstanceRow {
  createdAt: string;
  discountId: number;
  id: number;
  instanceInfo: string;
  productId: number;
  storeId: number;
  updatedAt: string;
}

export const DiscountInstanceRowTable = "discount_instances" as const;

export const DiscountInstanceRowColumns = {
  createdAt: "created_at",
  discountId: "discount_id",
  id: "id",
  instanceInfo: "instance_info",
  productId: "product_id",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export type DiscountInstanceRowColumn = keyof typeof DiscountInstanceRowColumns;

export const DiscountInstanceRowOrderBy = {
  createdAt: "created_at",
  discountId: "discount_id",
  id: "id",
  instanceInfo: "instance_info",
  productId: "product_id",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export interface DiscountRow {
  code: string;
  createdAt: string;
  currency: string;
  datetimeZone: string;
  description: string;
  duration: number;
  endDate: string;
  id: number;
  locale: string;
  maxAmount: number;
  minAmount: number;
  name: string;
  scope: string;
  startDate: string;
  status: string;
  target: string;
  type: string;
  updatedAt: string;
  value: number;
}

export const DiscountRowTable = "discounts" as const;

export const DiscountRowColumns = {
  code: "code",
  createdAt: "created_at",
  currency: "currency",
  datetimeZone: "datetime_zone",
  description: "description",
  duration: "duration",
  endDate: "end_date",
  id: "id",
  locale: "locale",
  maxAmount: "max_amount",
  minAmount: "min_amount",
  name: "name",
  scope: "scope",
  startDate: "start_date",
  status: "status",
  target: "target",
  type: "type",
  updatedAt: "updated_at",
  value: "value",
} as const;

export type DiscountRowColumn = keyof typeof DiscountRowColumns;

export const DiscountRowOrderBy = {
  code: "code",
  createdAt: "created_at",
  currency: "currency",
  datetimeZone: "datetime_zone",
  description: "description",
  duration: "duration",
  endDate: "end_date",
  id: "id",
  locale: "locale",
  maxAmount: "max_amount",
  minAmount: "min_amount",
  name: "name",
  scope: "scope",
  startDate: "start_date",
  status: "status",
  target: "target",
  type: "type",
  updatedAt: "updated_at",
  value: "value",
} as const;

export interface FailedJobRow {
  connection: string;
  exception: string;
  failedAt: string;
  id: number;
  payload: string;
  queue: string;
  uuid: string;
}

export const FailedJobRowTable = "failed_jobs" as const;

export const FailedJobRowColumns = {
  connection: "connection",
  exception: "exception",
  failedAt: "failed_at",
  id: "id",
  payload: "payload",
  queue: "queue",
  uuid: "uuid",
} as const;

export type FailedJobRowColumn = keyof typeof FailedJobRowColumns;

export const FailedJobRowOrderBy = {
  connection: "connection",
  exception: "exception",
  failedAt: "failed_at",
  id: "id",
  payload: "payload",
  queue: "queue",
  uuid: "uuid",
} as const;

export interface JobRow {
  attempts: number;
  availableAt: number;
  createdAt: number;
  id: number;
  payload: string;
  queue: string;
  reservedAt: number;
}

export const JobRowTable = "jobs" as const;

export const JobRowColumns = {
  attempts: "attempts",
  availableAt: "available_at",
  createdAt: "created_at",
  id: "id",
  payload: "payload",
  queue: "queue",
  reservedAt: "reserved_at",
} as const;

export type JobRowColumn = keyof typeof JobRowColumns;

export const JobRowOrderBy = {
  attempts: "attempts",
  availableAt: "available_at",
  createdAt: "created_at",
  id: "id",
  payload: "payload",
  queue: "queue",
  reservedAt: "reserved_at",
} as const;

export interface MenuRow {
  categoryIds: string;
  createdAt: string;
  description: string;
  id: number;
  name: string;
  schedule: string;
  status: string;
  storeId: number;
  updatedAt: string;
}

export const MenuRowTable = "menus" as const;

export const MenuRowColumns = {
  categoryIds: "category_ids",
  createdAt: "created_at",
  description: "description",
  id: "id",
  name: "name",
  schedule: "schedule",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export type MenuRowColumn = keyof typeof MenuRowColumns;

export const MenuRowOrderBy = {
  categoryIds: "category_ids",
  createdAt: "created_at",
  description: "description",
  id: "id",
  name: "name",
  schedule: "schedule",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
} as const;

export interface MigrationRow {
  batch: number;
  id: number;
  migration: string;
}

export const MigrationRowTable = "migrations" as const;

export const MigrationRowColumns = {
  batch: "batch",
  id: "id",
  migration: "migration",
} as const;

export type MigrationRowColumn = keyof typeof MigrationRowColumns;

export const MigrationRowOrderBy = {
  batch: "batch",
  id: "id",
  migration: "migration",
} as const;

export interface NotificationRow {
  body: string;
  createdAt: string;
  data: string;
  fromStoreId: number;
  fromUserId: number;
  id: number;
  name: string;
  notificationTime: string;
  orderId: string;
  readTime: string;
  retrievedFirstTime: string;
  source: string;
  title: string;
  toStoreId: number;
  toUserId: number;
  type: string;
  unread: string;
  updatedAt: string;
  url: string;
}

export const NotificationRowTable = "notifications" as const;

export const NotificationRowColumns = {
  body: "body",
  createdAt: "created_at",
  data: "data",
  fromStoreId: "from_store_Id",
  fromUserId: "from_user_id",
  id: "id",
  name: "name",
  notificationTime: "notification_time",
  orderId: "order_id",
  readTime: "read_time",
  retrievedFirstTime: "retrieved_first_time",
  source: "source",
  title: "title",
  toStoreId: "to_store_id",
  toUserId: "to_user_id",
  type: "type",
  unread: "unread",
  updatedAt: "updated_at",
  url: "url",
} as const;

export type NotificationRowColumn = keyof typeof NotificationRowColumns;

export const NotificationRowOrderBy = {
  body: "body",
  createdAt: "created_at",
  data: "data",
  fromStoreId: "from_store_Id",
  fromUserId: "from_user_id",
  id: "id",
  name: "name",
  notificationTime: "notification_time",
  orderId: "order_id",
  readTime: "read_time",
  retrievedFirstTime: "retrieved_first_time",
  source: "source",
  title: "title",
  toStoreId: "to_store_id",
  toUserId: "to_user_id",
  type: "type",
  unread: "unread",
  updatedAt: "updated_at",
  url: "url",
} as const;

export interface OrderProductRow {
  createdAt: string;
  id: number;
  orderId: number;
  productId: number;
  quantity: number;
  updatedAt: string;
}

export const OrderProductRowTable = "order_product" as const;

export const OrderProductRowColumns = {
  createdAt: "created_at",
  id: "id",
  orderId: "order_id",
  productId: "product_id",
  quantity: "quantity",
  updatedAt: "updated_at",
} as const;

export type OrderProductRowColumn = keyof typeof OrderProductRowColumns;

export const OrderProductRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  orderId: "order_id",
  productId: "product_id",
  quantity: "quantity",
  updatedAt: "updated_at",
} as const;

export interface OrderRow {
  courierUuid: string;
  createdAt: string;
  deliveryTaxAmount: number;
  discountAmount: number;
  id: number;
  model: string;
  notes: string;
  orderId: string;
  sellerTaxAmount: number;
  sellerTotal: number;
  status: string;
  storeId: number;
  subTotal: number;
  taxAmount: number;
  total: number;
  updatedAt: string;
  userId: number;
}

export const OrderRowTable = "orders" as const;

export const OrderRowColumns = {
  courierUuid: "courier_uuid",
  createdAt: "created_at",
  deliveryTaxAmount: "delivery_tax_amount",
  discountAmount: "discount_amount",
  id: "id",
  model: "model",
  notes: "notes",
  orderId: "order_id",
  sellerTaxAmount: "seller_tax_amount",
  sellerTotal: "seller_total",
  status: "status",
  storeId: "store_id",
  subTotal: "sub_total",
  taxAmount: "tax_amount",
  total: "total",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type OrderRowColumn = keyof typeof OrderRowColumns;

export const OrderRowOrderBy = {
  courierUuid: "courier_uuid",
  createdAt: "created_at",
  deliveryTaxAmount: "delivery_tax_amount",
  discountAmount: "discount_amount",
  id: "id",
  model: "model",
  notes: "notes",
  orderId: "order_id",
  sellerTaxAmount: "seller_tax_amount",
  sellerTotal: "seller_total",
  status: "status",
  storeId: "store_id",
  subTotal: "sub_total",
  taxAmount: "tax_amount",
  total: "total",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface OrgOrderRow {
  billingAddress: string;
  billingCity: string;
  billingDiscount: number;
  billingDiscountCode: string;
  billingEmail: string;
  billingName: string;
  billingNameOnCard: string;
  billingPhone: string;
  billingPostalcode: string;
  billingProvince: string;
  billingSubtotal: number;
  billingTax: number;
  billingTotal: number;
  createdAt: string;
  error: string;
  id: number;
  paymentGateway: string;
  shipped: number;
  updatedAt: string;
  userId: number;
}

export const OrgOrderRowTable = "org_orders" as const;

export const OrgOrderRowColumns = {
  billingAddress: "billing_address",
  billingCity: "billing_city",
  billingDiscount: "billing_discount",
  billingDiscountCode: "billing_discount_code",
  billingEmail: "billing_email",
  billingName: "billing_name",
  billingNameOnCard: "billing_name_on_card",
  billingPhone: "billing_phone",
  billingPostalcode: "billing_postalcode",
  billingProvince: "billing_province",
  billingSubtotal: "billing_subtotal",
  billingTax: "billing_tax",
  billingTotal: "billing_total",
  createdAt: "created_at",
  error: "error",
  id: "id",
  paymentGateway: "payment_gateway",
  shipped: "shipped",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type OrgOrderRowColumn = keyof typeof OrgOrderRowColumns;

export const OrgOrderRowOrderBy = {
  billingAddress: "billing_address",
  billingCity: "billing_city",
  billingDiscount: "billing_discount",
  billingDiscountCode: "billing_discount_code",
  billingEmail: "billing_email",
  billingName: "billing_name",
  billingNameOnCard: "billing_name_on_card",
  billingPhone: "billing_phone",
  billingPostalcode: "billing_postalcode",
  billingProvince: "billing_province",
  billingSubtotal: "billing_subtotal",
  billingTax: "billing_tax",
  billingTotal: "billing_total",
  createdAt: "created_at",
  error: "error",
  id: "id",
  paymentGateway: "payment_gateway",
  shipped: "shipped",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface PasswordResetRow {
  createdAt: string;
  email: string;
  token: string;
}

export const PasswordResetRowTable = "password_resets" as const;

export const PasswordResetRowColumns = {
  createdAt: "created_at",
  email: "email",
  token: "token",
} as const;

export type PasswordResetRowColumn = keyof typeof PasswordResetRowColumns;

export const PasswordResetRowOrderBy = {
  createdAt: "created_at",
  email: "email",
  token: "token",
} as const;

export interface PersonalAccessTokenRow {
  abilities: string;
  createdAt: string;
  id: number;
  lastUsedAt: string;
  name: string;
  token: string;
  tokenableId: number;
  tokenableType: string;
  updatedAt: string;
}

export const PersonalAccessTokenRowTable = "personal_access_tokens" as const;

export const PersonalAccessTokenRowColumns = {
  abilities: "abilities",
  createdAt: "created_at",
  id: "id",
  lastUsedAt: "last_used_at",
  name: "name",
  token: "token",
  tokenableId: "tokenable_id",
  tokenableType: "tokenable_type",
  updatedAt: "updated_at",
} as const;

export type PersonalAccessTokenRowColumn = keyof typeof PersonalAccessTokenRowColumns;

export const PersonalAccessTokenRowOrderBy = {
  abilities: "abilities",
  createdAt: "created_at",
  id: "id",
  lastUsedAt: "last_used_at",
  name: "name",
  token: "token",
  tokenableId: "tokenable_id",
  tokenableType: "tokenable_type",
  updatedAt: "updated_at",
} as const;

export interface ProductCustomizationRelationRow {
  createdAt: string;
  id: number;
  parentId: number;
  position: number;
  productCustomizationId: number;
  productId: number;
  updatedAt: string;
  userId: number;
}

export const ProductCustomizationRelationRowTable = "product_customization_relations" as const;

export const ProductCustomizationRelationRowColumns = {
  createdAt: "created_at",
  id: "id",
  parentId: "parent_id",
  position: "position",
  productCustomizationId: "product_customization_id",
  productId: "product_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type ProductCustomizationRelationRowColumn = keyof typeof ProductCustomizationRelationRowColumns;

export const ProductCustomizationRelationRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  parentId: "parent_id",
  position: "position",
  productCustomizationId: "product_customization_id",
  productId: "product_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface ProductCustomizationRow {
  createdAt: string;
  defaultValue: string;
  description: string;
  id: number;
  isOption: string;
  isSoldOut: string;
  maxQuantity: number;
  minQuantity: number;
  name: string;
  nutritionalInfo: string;
  pickupPrice: number;
  price: number;
  showInOrder: string;
  status: string;
  type: string;
  updatedAt: string;
  userId: number;
  uuid: string;
}

export const ProductCustomizationRowTable = "product_customizations" as const;

export const ProductCustomizationRowColumns = {
  createdAt: "created_at",
  defaultValue: "default_value",
  description: "description",
  id: "id",
  isOption: "is_option",
  isSoldOut: "is_sold_out",
  maxQuantity: "max_quantity",
  minQuantity: "min_quantity",
  name: "name",
  nutritionalInfo: "nutritional_info",
  pickupPrice: "pickup_price",
  price: "price",
  showInOrder: "show_in_order",
  status: "status",
  type: "type",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export type ProductCustomizationRowColumn = keyof typeof ProductCustomizationRowColumns;

export const ProductCustomizationRowOrderBy = {
  createdAt: "created_at",
  defaultValue: "default_value",
  description: "description",
  id: "id",
  isOption: "is_option",
  isSoldOut: "is_sold_out",
  maxQuantity: "max_quantity",
  minQuantity: "min_quantity",
  name: "name",
  nutritionalInfo: "nutritional_info",
  pickupPrice: "pickup_price",
  price: "price",
  showInOrder: "show_in_order",
  status: "status",
  type: "type",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export interface ProductOrderRow {
  createdAt: string;
  id: number;
  position: number;
  productId: number;
  updatedAt: string;
  userId: number;
}

export const ProductOrderRowTable = "product_order" as const;

export const ProductOrderRowColumns = {
  createdAt: "created_at",
  id: "id",
  position: "position",
  productId: "product_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type ProductOrderRowColumn = keyof typeof ProductOrderRowColumns;

export const ProductOrderRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  position: "position",
  productId: "product_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface ProductRow {
  createdAt: string;
  deliveryPrice: number;
  depositOptions: string;
  description: string;
  details: string;
  featured: number;
  id: number;
  images: string;
  keywords: string;
  name: string;
  price: number;
  quantity: number;
  slug: string;
  status: string;
  storeId: number;
  updatedAt: string;
  userId: number;
  uuid: string;
}

export const ProductRowTable = "products" as const;

export const ProductRowColumns = {
  createdAt: "created_at",
  deliveryPrice: "delivery_price",
  depositOptions: "deposit_options",
  description: "description",
  details: "details",
  featured: "featured",
  id: "id",
  images: "images",
  keywords: "keywords",
  name: "name",
  price: "price",
  quantity: "quantity",
  slug: "slug",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export type ProductRowColumn = keyof typeof ProductRowColumns;

export const ProductRowOrderBy = {
  createdAt: "created_at",
  deliveryPrice: "delivery_price",
  depositOptions: "deposit_options",
  description: "description",
  details: "details",
  featured: "featured",
  id: "id",
  images: "images",
  keywords: "keywords",
  name: "name",
  price: "price",
  quantity: "quantity",
  slug: "slug",
  status: "status",
  storeId: "store_id",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export interface SessionRow {
  id: string;
  ipAddress: string;
  lastActivity: number;
  payload: string;
  userAgent: string;
  userId: number;
}

export const SessionRowTable = "sessions" as const;

export const SessionRowColumns = {
  id: "id",
  ipAddress: "ip_address",
  lastActivity: "last_activity",
  payload: "payload",
  userAgent: "user_agent",
  userId: "user_id",
} as const;

export type SessionRowColumn = keyof typeof SessionRowColumns;

export const SessionRowOrderBy = {
  id: "id",
  ipAddress: "ip_address",
  lastActivity: "last_activity",
  payload: "payload",
  userAgent: "user_agent",
  userId: "user_id",
} as const;

export interface ShoppingcartRow {
  content: string;
  createdAt: string;
  identifier: string;
  instance: string;
  updatedAt: string;
}

export const ShoppingcartRowTable = "shoppingcart" as const;

export const ShoppingcartRowColumns = {
  content: "content",
  createdAt: "created_at",
  identifier: "identifier",
  instance: "instance",
  updatedAt: "updated_at",
} as const;

export type ShoppingcartRowColumn = keyof typeof ShoppingcartRowColumns;

export const ShoppingcartRowOrderBy = {
  content: "content",
  createdAt: "created_at",
  identifier: "identifier",
  instance: "instance",
  updatedAt: "updated_at",
} as const;

export interface StoreRow {
  businessHours: string;
  businessPhone: string;
  businessPhoto: string;
  categories: string;
  createdAt: string;
  currency: string;
  datetimeZone: string;
  deliveryAddressId: number;
  deliveryCost: number;
  deliveryOptions: string;
  deliveryTimeframeOptions: string;
  id: number;
  isAcceptingOrders: string;
  isOpen: string;
  isPublished: string;
  locale: string;
  name: string;
  orderStateConfig: string;
  paymentOptions: string;
  pickupAddressId: number;
  serviceFeeOptions: string;
  slogan: string;
  slug: string;
  status: string;
  updatedAt: string;
  userId: number;
  uuid: string;
}

export const StoreRowTable = "stores" as const;

export const StoreRowColumns = {
  businessHours: "business_hours",
  businessPhone: "business_phone",
  businessPhoto: "business_photo",
  categories: "categories",
  createdAt: "created_at",
  currency: "currency",
  datetimeZone: "datetime_zone",
  deliveryAddressId: "delivery_address_id",
  deliveryCost: "delivery_cost",
  deliveryOptions: "delivery_options",
  deliveryTimeframeOptions: "delivery_timeframe_options",
  id: "id",
  isAcceptingOrders: "is_accepting_orders",
  isOpen: "is_open",
  isPublished: "is_published",
  locale: "locale",
  name: "name",
  orderStateConfig: "order_state_config",
  paymentOptions: "payment_options",
  pickupAddressId: "pickup_address_id",
  serviceFeeOptions: "service_fee_options",
  slogan: "slogan",
  slug: "slug",
  status: "status",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export type StoreRowColumn = keyof typeof StoreRowColumns;

export const StoreRowOrderBy = {
  businessHours: "business_hours",
  businessPhone: "business_phone",
  businessPhoto: "business_photo",
  categories: "categories",
  createdAt: "created_at",
  currency: "currency",
  datetimeZone: "datetime_zone",
  deliveryAddressId: "delivery_address_id",
  deliveryCost: "delivery_cost",
  deliveryOptions: "delivery_options",
  deliveryTimeframeOptions: "delivery_timeframe_options",
  id: "id",
  isAcceptingOrders: "is_accepting_orders",
  isOpen: "is_open",
  isPublished: "is_published",
  locale: "locale",
  name: "name",
  orderStateConfig: "order_state_config",
  paymentOptions: "payment_options",
  pickupAddressId: "pickup_address_id",
  serviceFeeOptions: "service_fee_options",
  slogan: "slogan",
  slug: "slug",
  status: "status",
  updatedAt: "updated_at",
  userId: "user_id",
  uuid: "uuid",
} as const;

export interface TeamUserRow {
  createdAt: string;
  id: number;
  role: string;
  teamId: number;
  updatedAt: string;
  userId: number;
}

export const TeamUserRowTable = "team_user" as const;

export const TeamUserRowColumns = {
  createdAt: "created_at",
  id: "id",
  role: "role",
  teamId: "team_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type TeamUserRowColumn = keyof typeof TeamUserRowColumns;

export const TeamUserRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  role: "role",
  teamId: "team_id",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface TeamRow {
  createdAt: string;
  id: number;
  name: string;
  personalTeam: number;
  updatedAt: string;
  userId: number;
}

export const TeamRowTable = "teams" as const;

export const TeamRowColumns = {
  createdAt: "created_at",
  id: "id",
  name: "name",
  personalTeam: "personal_team",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export type TeamRowColumn = keyof typeof TeamRowColumns;

export const TeamRowOrderBy = {
  createdAt: "created_at",
  id: "id",
  name: "name",
  personalTeam: "personal_team",
  updatedAt: "updated_at",
  userId: "user_id",
} as const;

export interface UserRow {
  canSell: string;
  cashAppId: string;
  createdAt: string;
  currentTeamId: number;
  datetimeZone: string;
  deliveryAddressId: number;
  email: string;
  emailVerifiedAt: string;
  facebookId: string;
  googleId: string;
  id: number;
  isAvailable: string;
  isSystemAdmin: string;
  isTestUser: string;
  lastLocation: string;
  lastLocationTime: string;
  locale: string;
  name: string;
  password: string;
  phone: string;
  profilePhotoPath: string;
  rememberToken: string;
  role: string;
  status: string;
  storeId: number;
  testSeller: string;
  twoFactorRecoveryCodes: string;
  twoFactorSecret: string;
  updatedAt: string;
  uuid: string;
  zelleEmail: string;
  zellePhone: string;
}

export const UserRowTable = "users" as const;

export const UserRowColumns = {
  canSell: "can_sell",
  cashAppId: "cash_app_id",
  createdAt: "created_at",
  currentTeamId: "current_team_id",
  datetimeZone: "datetime_zone",
  deliveryAddressId: "delivery_address_id",
  email: "email",
  emailVerifiedAt: "email_verified_at",
  facebookId: "facebook_id",
  googleId: "google_id",
  id: "id",
  isAvailable: "is_available",
  isSystemAdmin: "is_system_admin",
  isTestUser: "is_test_user",
  lastLocation: "last_location",
  lastLocationTime: "last_location_time",
  locale: "locale",
  name: "name",
  password: "password",
  phone: "phone",
  profilePhotoPath: "profile_photo_path",
  rememberToken: "remember_token",
  role: "role",
  status: "status",
  storeId: "store_id",
  testSeller: "test_seller",
  twoFactorRecoveryCodes: "two_factor_recovery_codes",
  twoFactorSecret: "two_factor_secret",
  updatedAt: "updated_at",
  uuid: "uuid",
  zelleEmail: "zelle_email",
  zellePhone: "zelle_phone",
} as const;

export type UserRowColumn = keyof typeof UserRowColumns;

export const UserRowOrderBy = {
  canSell: "can_sell",
  cashAppId: "cash_app_id",
  createdAt: "created_at",
  currentTeamId: "current_team_id",
  datetimeZone: "datetime_zone",
  deliveryAddressId: "delivery_address_id",
  email: "email",
  emailVerifiedAt: "email_verified_at",
  facebookId: "facebook_id",
  googleId: "google_id",
  id: "id",
  isAvailable: "is_available",
  isSystemAdmin: "is_system_admin",
  isTestUser: "is_test_user",
  lastLocation: "last_location",
  lastLocationTime: "last_location_time",
  locale: "locale",
  name: "name",
  password: "password",
  phone: "phone",
  profilePhotoPath: "profile_photo_path",
  rememberToken: "remember_token",
  role: "role",
  status: "status",
  storeId: "store_id",
  testSeller: "test_seller",
  twoFactorRecoveryCodes: "two_factor_recovery_codes",
  twoFactorSecret: "two_factor_secret",
  updatedAt: "updated_at",
  uuid: "uuid",
  zelleEmail: "zelle_email",
  zellePhone: "zelle_phone",
} as const;

export interface WebsocketsStatisticsEntryRow {
  apiMessageCount: number;
  appId: string;
  createdAt: string;
  id: number;
  peakConnectionCount: number;
  updatedAt: string;
  websocketMessageCount: number;
}

export const WebsocketsStatisticsEntryRowTable = "websockets_statistics_entries" as const;

export const WebsocketsStatisticsEntryRowColumns = {
  apiMessageCount: "api_message_count",
  appId: "app_id",
  createdAt: "created_at",
  id: "id",
  peakConnectionCount: "peak_connection_count",
  updatedAt: "updated_at",
  websocketMessageCount: "websocket_message_count",
} as const;

export type WebsocketsStatisticsEntryRowColumn = keyof typeof WebsocketsStatisticsEntryRowColumns;

export const WebsocketsStatisticsEntryRowOrderBy = {
  apiMessageCount: "api_message_count",
  appId: "app_id",
  createdAt: "created_at",
  id: "id",
  peakConnectionCount: "peak_connection_count",
  updatedAt: "updated_at",
  websocketMessageCount: "websocket_message_count",
} as const;


