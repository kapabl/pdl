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
 * @class MenuRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $categoryIds
 * @property string $createdAt
 * @property string $description
 * @property int $id
 * @property string $name
 * @property string $schedule
 * @property string $status
 * @property int $storeId
 * @property string $updatedAt
 *
 */
class MenuRow extends Row
{

    const TableName = 'menus';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.menus';

    /**
     * MenuRow constructor.
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
     * @return MenuWhere
     */
    public static function where()
    {
        $result = new MenuWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return MenuFieldList
     */
    public static function fieldList()
    {
        $result = new MenuFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

