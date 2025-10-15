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
 * @class PasswordResetRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property string $email
 * @property string $token
 *
 */
class PasswordResetRow extends Row
{

    const TableName = 'password_resets';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.password_resets';

    /**
     * PasswordResetRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];

        parent::__construct( ...$arguments );
    }

    /**
     * @return PasswordResetWhere
     */
    public static function where()
    {
        $result = new PasswordResetWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return PasswordResetFieldList
     */
    public static function fieldList()
    {
        $result = new PasswordResetFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

