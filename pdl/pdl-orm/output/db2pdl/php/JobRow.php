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
 * @class JobRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property int $attempts
 * @property int $availableAt
 * @property int $createdAt
 * @property int $id
 * @property string $payload
 * @property string $queue
 * @property int $reservedAt
 *
 */
class JobRow extends Row
{

    const TableName = 'jobs';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.jobs';

    /**
     * JobRow constructor.
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
     * @return JobWhere
     */
    public static function where()
    {
        $result = new JobWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return JobFieldList
     */
    public static function fieldList()
    {
        $result = new JobFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

