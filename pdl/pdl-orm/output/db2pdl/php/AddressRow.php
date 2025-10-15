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
 * @class AddressRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $address1
 * @property string $address2
 * @property string $city
 * @property string $country
 * @property string $createdAt
 * @property string $defaultDelivery
 * @property string $defaultPickup
 * @property int $id
 * @property string $isTest
 * @property float $lat
 * @property float $lon
 * @property string $name
 * @property string $phone
 * @property string $state
 * @property string $status
 * @property string $updatedAt
 * @property int $userId
 * @property string $zipcode
 *
 */
class AddressRow extends Row
{

    const TableName = 'addresses';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.addresses';

    /**
     * AddressRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];
        $this->_propertyAttributes[ 'address1' ] = [ 'ColumnName' => [ 'default1' => '{}' ] ];
        $this->_propertyAttributes[ 'address2' ] = [ 'ColumnName' => [ 'default1' => '{}' ] ];
        $this->_propertyAttributes[ 'id' ] = [ 'IsDbId' => [] ];

        parent::__construct( ...$arguments );
    }

    /**
     * @return AddressWhere
     */
    public static function where()
    {
        $result = new AddressWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return AddressFieldList
     */
    public static function fieldList()
    {
        $result = new AddressFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

