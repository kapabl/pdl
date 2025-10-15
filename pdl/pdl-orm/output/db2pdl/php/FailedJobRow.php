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
 * @class FailedJobRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $connection
 * @property string $exception
 * @property string $failedAt
 * @property int $id
 * @property string $payload
 * @property string $queue
 * @property string $uuid
 *
 */
class FailedJobRow extends Row
{

    const TableName = 'failed_jobs';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.failed_jobs';

    /**
     * FailedJobRow constructor.
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
     * @return FailedJobWhere
     */
    public static function where()
    {
        $result = new FailedJobWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return FailedJobFieldList
     */
    public static function fieldList()
    {
        $result = new FailedJobFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

