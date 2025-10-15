<?php
/**
 *  Minglehouse Generated File with Pdl
 *  MiManjar
 *  @version: 1.0.0
 *
 */

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Row;
/**
 * @class OrgOrderRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $billingAddress
 * @property string $billingCity
 * @property int $billingDiscount
 * @property string $billingDiscountCode
 * @property string $billingEmail
 * @property string $billingName
 * @property string $billingNameOnCard
 * @property string $billingPhone
 * @property string $billingPostalcode
 * @property string $billingProvince
 * @property int $billingSubtotal
 * @property int $billingTax
 * @property int $billingTotal
 * @property string $createdAt
 * @property string $error
 * @property int $id
 * @property string $paymentGateway
 * @property int $shipped
 * @property string $updatedAt
 * @property int $userId
 *
 */
class OrgOrderRow extends Row
{

    const TableName = 'org_orders';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.org_orders';

    /**
     * OrgOrderRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];
        $this->_propertyAttributes[ 'id' ] = [ 'IsDbId' => [] ];

        parent::__construct( ...$arguments );
    }

    /**
     * @return OrgOrderWhere
     */
    public static function where()
    {
        $result = new OrgOrderWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return OrgOrderFieldList
     */
    public static function fieldList()
    {
        $result = new OrgOrderFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

