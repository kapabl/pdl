<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class UserColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'can_sell', 'string' );
        $this->addColumn( 'cash_app_id', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'current_team_id', 'int' );
        $this->addColumn( 'datetime_zone', 'string' );
        $this->addColumn( 'delivery_address_id', 'int' );
        $this->addColumn( 'email', 'string' );
        $this->addColumn( 'email_verified_at', 'string' );
        $this->addColumn( 'facebook_id', 'string' );
        $this->addColumn( 'google_id', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'is_available', 'string' );
        $this->addColumn( 'is_system_admin', 'string' );
        $this->addColumn( 'is_test_user', 'string' );
        $this->addColumn( 'last_location', 'string' );
        $this->addColumn( 'last_location_time', 'string' );
        $this->addColumn( 'locale', 'string' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'password', 'string' );
        $this->addColumn( 'phone', 'string' );
        $this->addColumn( 'profile_photo_path', 'string' );
        $this->addColumn( 'remember_token', 'string' );
        $this->addColumn( 'role', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'store_id', 'int' );
        $this->addColumn( 'test_seller', 'string' );
        $this->addColumn( 'two_factor_recovery_codes', 'string' );
        $this->addColumn( 'two_factor_secret', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'uuid', 'string' );
        $this->addColumn( 'zelle_email', 'string' );
        $this->addColumn( 'zelle_phone', 'string' );
    }
}

