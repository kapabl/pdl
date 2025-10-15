<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait UserColumnsTraits
{
    public function canSell()
    {
        $this->addColumn( 'can_sell' );
        return $this;
    }

    public function cashAppId()
    {
        $this->addColumn( 'cash_app_id' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function currentTeamId()
    {
        $this->addColumn( 'current_team_id' );
        return $this;
    }

    public function datetimeZone()
    {
        $this->addColumn( 'datetime_zone' );
        return $this;
    }

    public function deliveryAddressId()
    {
        $this->addColumn( 'delivery_address_id' );
        return $this;
    }

    public function email()
    {
        $this->addColumn( 'email' );
        return $this;
    }

    public function emailVerifiedAt()
    {
        $this->addColumn( 'email_verified_at' );
        return $this;
    }

    public function facebookId()
    {
        $this->addColumn( 'facebook_id' );
        return $this;
    }

    public function googleId()
    {
        $this->addColumn( 'google_id' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function isAvailable()
    {
        $this->addColumn( 'is_available' );
        return $this;
    }

    public function isSystemAdmin()
    {
        $this->addColumn( 'is_system_admin' );
        return $this;
    }

    public function isTestUser()
    {
        $this->addColumn( 'is_test_user' );
        return $this;
    }

    public function lastLocation()
    {
        $this->addColumn( 'last_location' );
        return $this;
    }

    public function lastLocationTime()
    {
        $this->addColumn( 'last_location_time' );
        return $this;
    }

    public function locale()
    {
        $this->addColumn( 'locale' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function password()
    {
        $this->addColumn( 'password' );
        return $this;
    }

    public function phone()
    {
        $this->addColumn( 'phone' );
        return $this;
    }

    public function profilePhotoPath()
    {
        $this->addColumn( 'profile_photo_path' );
        return $this;
    }

    public function rememberToken()
    {
        $this->addColumn( 'remember_token' );
        return $this;
    }

    public function role()
    {
        $this->addColumn( 'role' );
        return $this;
    }

    public function status()
    {
        $this->addColumn( 'status' );
        return $this;
    }

    public function storeId()
    {
        $this->addColumn( 'store_id' );
        return $this;
    }

    public function testSeller()
    {
        $this->addColumn( 'test_seller' );
        return $this;
    }

    public function twoFactorRecoveryCodes()
    {
        $this->addColumn( 'two_factor_recovery_codes' );
        return $this;
    }

    public function twoFactorSecret()
    {
        $this->addColumn( 'two_factor_secret' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function uuid()
    {
        $this->addColumn( 'uuid' );
        return $this;
    }

    public function zelleEmail()
    {
        $this->addColumn( 'zelle_email' );
        return $this;
    }

    public function zellePhone()
    {
        $this->addColumn( 'zelle_phone' );
        return $this;
    }

}

