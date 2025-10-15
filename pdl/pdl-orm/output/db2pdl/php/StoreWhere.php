<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class StoreWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'business_hours' );
        $this->addField( 'business_phone' );
        $this->addField( 'business_photo' );
        $this->addField( 'categories' );
        $this->addField( 'created_at' );
        $this->addField( 'currency' );
        $this->addField( 'datetime_zone' );
        $this->addField( 'delivery_address_id' );
        $this->addField( 'delivery_cost' );
        $this->addField( 'delivery_options' );
        $this->addField( 'delivery_timeframe_options' );
        $this->addField( 'id' );
        $this->addField( 'is_accepting_orders' );
        $this->addField( 'is_open' );
        $this->addField( 'is_published' );
        $this->addField( 'locale' );
        $this->addField( 'name' );
        $this->addField( 'order_state_config' );
        $this->addField( 'payment_options' );
        $this->addField( 'pickup_address_id' );
        $this->addField( 'service_fee_options' );
        $this->addField( 'slogan' );
        $this->addField( 'slug' );
        $this->addField( 'status' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
        $this->addField( 'uuid' );
    }
}

