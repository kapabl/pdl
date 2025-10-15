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
 * @class UserRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $canSell
 * @property string $cashAppId
 * @property string $createdAt
 * @property int $currentTeamId
 * @property string $datetimeZone
 * @property int $deliveryAddressId
 * @property string $email
 * @property string $emailVerifiedAt
 * @property string $facebookId
 * @property string $googleId
 * @property int $id
 * @property string $isAvailable
 * @property string $isSystemAdmin
 * @property string $isTestUser
 * @property string $lastLocation
 * @property string $lastLocationTime
 * @property string $locale
 * @property string $name
 * @property string $password
 * @property string $phone
 * @property string $profilePhotoPath
 * @property string $rememberToken
 * @property string $role
 * @property string $status
 * @property int $storeId
 * @property string $testSeller
 * @property string $twoFactorRecoveryCodes
 * @property string $twoFactorSecret
 * @property string $updatedAt
 * @property string $uuid
 * @property string $zelleEmail
 * @property string $zellePhone
 *
 */
class UserRow extends Row
{

    const TableName = 'users';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.users';

    /**
     * UserRow constructor.
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
     * @return UserWhere
     */
    public static function where()
    {
        $result = new UserWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return UserFieldList
     */
    public static function fieldList()
    {
        $result = new UserFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

