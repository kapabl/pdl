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
 * @class CustomizationSearchRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property string $customizationFullName
 * @property int $id
 * @property string $name
 * @property int $relationId
 * @property string $updatedAt
 * @property int $userId
 *
 */
class CustomizationSearchRow extends Row
{

    const TableName = 'customization_search';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.customization_search';

    /**
     * CustomizationSearchRow constructor.
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
     * @return CustomizationSearchWhere
     */
    public static function where()
    {
        $result = new CustomizationSearchWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CustomizationSearchFieldList
     */
    public static function fieldList()
    {
        $result = new CustomizationSearchFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

