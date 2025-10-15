<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class StoreColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'business_hours', 'string' );
        $this->addColumn( 'business_phone', 'string' );
        $this->addColumn( 'business_photo', 'string' );
        $this->addColumn( 'categories', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'currency', 'string' );
        $this->addColumn( 'datetime_zone', 'string' );
        $this->addColumn( 'delivery_address_id', 'int' );
        $this->addColumn( 'delivery_cost', 'int' );
        $this->addColumn( 'delivery_options', 'string' );
        $this->addColumn( 'delivery_timeframe_options', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'is_accepting_orders', 'string' );
        $this->addColumn( 'is_open', 'string' );
        $this->addColumn( 'is_published', 'string' );
        $this->addColumn( 'locale', 'string' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'order_state_config', 'string' );
        $this->addColumn( 'payment_options', 'string' );
        $this->addColumn( 'pickup_address_id', 'int' );
        $this->addColumn( 'service_fee_options', 'string' );
        $this->addColumn( 'slogan', 'string' );
        $this->addColumn( 'slug', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
        $this->addColumn( 'uuid', 'string' );
    }
}

