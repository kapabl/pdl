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
 * @class WebsocketsStatisticsEntryRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property int $apiMessageCount
 * @property string $appId
 * @property string $createdAt
 * @property int $id
 * @property int $peakConnectionCount
 * @property string $updatedAt
 * @property int $websocketMessageCount
 *
 */
class WebsocketsStatisticsEntryRow extends Row
{

    const TableName = 'websockets_statistics_entries';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.websockets_statistics_entries';

    /**
     * WebsocketsStatisticsEntryRow constructor.
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
     * @return WebsocketsStatisticsEntryWhere
     */
    public static function where()
    {
        $result = new WebsocketsStatisticsEntryWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return WebsocketsStatisticsEntryFieldList
     */
    public static function fieldList()
    {
        $result = new WebsocketsStatisticsEntryFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

