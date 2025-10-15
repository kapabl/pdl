<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class DiscountColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'code', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'currency', 'string' );
        $this->addColumn( 'datetime_zone', 'string' );
        $this->addColumn( 'description', 'string' );
        $this->addColumn( 'duration', 'int' );
        $this->addColumn( 'end_date', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'locale', 'string' );
        $this->addColumn( 'max_amount', 'float' );
        $this->addColumn( 'min_amount', 'float' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'scope', 'string' );
        $this->addColumn( 'start_date', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'target', 'string' );
        $this->addColumn( 'type', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'value', 'float' );
    }
}

