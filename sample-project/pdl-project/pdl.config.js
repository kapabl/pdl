
let mergeConfig = require('./config/mergeConfig');

const Db2PdlSourceDest = "com/mh/mimanjar/domain/data";

const config = mergeConfig( {
    companyName: "Minglehouse",
    project: "MiManjar",
    version: "1.0.0",
    db2PdlSourceDest: Db2PdlSourceDest,
    sections: [
        {
            name: 'ConstFiles',

            files: {
                phpJsConsts: [
                    'com/mh/mimanjar/domain/products/CustomizationTypes.pdl',
                    'com/mh/mimanjar/domain/products/ProductCalculatedColumns.pdl',
                    'com/mh/mimanjar/domain/seller/DeliveryMethods.pdl',
                    'com/mh/mimanjar/domain/seller/StoreCalculatedColumns.pdl',
                    'com/mh/mimanjar/domain/seller/DeliveryTimeTypes.pdl',
                    'com/mh/mimanjar/domain/seller/PaymentMethods.pdl',
                    'com/mh/mimanjar/domain/seller/store/StoreEvents.pdl',
                    'com/mh/mimanjar/domain/TimeframeTypes.pdl',
                    'com/mh/mimanjar/domain/user/Roles.pdl',
                    'com/mh/mimanjar/domain/courier/TravelModes.pdl',
                    'com/mh/mimanjar/domain/courier/CourierOrderStatus.pdl',
                    'com/mh/mimanjar/domain/courier/CourierCalculatedColumns.pdl',
                    'com/mh/mimanjar/domain/order/OrderState.pdl',
                    'com/mh/mimanjar/domain/order/OrderTransitions.pdl',
                    'com/mh/mimanjar/domain/order/OrderSource.pdl',
                    'com/mh/mimanjar/domain/order/PaymentTimes.pdl',
                    'com/mh/mimanjar/domain/order/ActionReasonTypes.pdl',
                    'com/mh/mimanjar/domain/order/UserCancelReasonTypes.pdl',
                    'com/mh/mimanjar/domain/order/StoreCancelReasonTypes.pdl',
                    'com/mh/mimanjar/domain/order/OrderEventSource.pdl',
                    'com/mh/mimanjar/domain/order/TaxCalculationMethods.pdl',
                    'com/mh/mimanjar/domain/categories/CategoryCalculatedColumns.pdl',
                    'com/mh/mimanjar/domain/discounts/DiscountScopes.pdl',
                    'com/mh/mimanjar/domain/discounts/DiscountTypes.pdl',
                    'com/mh/mimanjar/domain/discounts/DiscountTargets.pdl',
                    'com/mh/mimanjar/domain/discounts/DiscountCalculatedColumns.pdl',
                    'com/mh/mimanjar/domain/discounts/DiscountInstanceCalculatedColumns.pdl'
                ],
                phpJsConstsExclude: [
                    Db2PdlSourceDest + '/*.pdl'
                ]
            }
        }
    ],
});

module.exports = config;


