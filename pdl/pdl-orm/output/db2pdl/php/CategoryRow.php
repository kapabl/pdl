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
 * @class CategoryRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property string $name
 * @property int $position
 * @property string $slug
 * @property string $status
 * @property int $storeId
 * @property string $updatedAt
 *
 */
class CategoryRow extends Row
{

    const TableName = 'categories';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.categories';

    /**
     * CategoryRow constructor.
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
     * @return CategoryWhere
     */
    public static function where()
    {
        $result = new CategoryWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CategoryFieldList
     */
    public static function fieldList()
    {
        $result = new CategoryFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

