<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class DiscountFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'code' );
        $this->addField( 'created_at' );
        $this->addField( 'currency' );
        $this->addField( 'datetime_zone' );
        $this->addField( 'description' );
        $this->addField( 'duration' );
        $this->addField( 'end_date' );
        $this->addField( 'id' );
        $this->addField( 'locale' );
        $this->addField( 'max_amount' );
        $this->addField( 'min_amount' );
        $this->addField( 'name' );
        $this->addField( 'scope' );
        $this->addField( 'start_date' );
        $this->addField( 'status' );
        $this->addField( 'target' );
        $this->addField( 'type' );
        $this->addField( 'updated_at' );
        $this->addField( 'value' );
    }
}

