<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class ProductCustomizationOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'default_value' );
        $this->addField( 'description' );
        $this->addField( 'id' );
        $this->addField( 'is_option' );
        $this->addField( 'is_sold_out' );
        $this->addField( 'max_quantity' );
        $this->addField( 'min_quantity' );
        $this->addField( 'name' );
        $this->addField( 'nutritional_info' );
        $this->addField( 'pickup_price' );
        $this->addField( 'price' );
        $this->addField( 'show_in_order' );
        $this->addField( 'status' );
        $this->addField( 'type' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
        $this->addField( 'uuid' );
    }
}

