<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class UserWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'can_sell' );
        $this->addField( 'cash_app_id' );
        $this->addField( 'created_at' );
        $this->addField( 'current_team_id' );
        $this->addField( 'datetime_zone' );
        $this->addField( 'delivery_address_id' );
        $this->addField( 'email' );
        $this->addField( 'email_verified_at' );
        $this->addField( 'facebook_id' );
        $this->addField( 'google_id' );
        $this->addField( 'id' );
        $this->addField( 'is_available' );
        $this->addField( 'is_system_admin' );
        $this->addField( 'is_test_user' );
        $this->addField( 'last_location' );
        $this->addField( 'last_location_time' );
        $this->addField( 'locale' );
        $this->addField( 'name' );
        $this->addField( 'password' );
        $this->addField( 'phone' );
        $this->addField( 'profile_photo_path' );
        $this->addField( 'remember_token' );
        $this->addField( 'role' );
        $this->addField( 'status' );
        $this->addField( 'store_id' );
        $this->addField( 'test_seller' );
        $this->addField( 'two_factor_recovery_codes' );
        $this->addField( 'two_factor_secret' );
        $this->addField( 'updated_at' );
        $this->addField( 'uuid' );
        $this->addField( 'zelle_email' );
        $this->addField( 'zelle_phone' );
    }
}

