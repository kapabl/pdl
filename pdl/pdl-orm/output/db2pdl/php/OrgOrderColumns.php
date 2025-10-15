<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class OrgOrderColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'billing_address', 'string' );
        $this->addColumn( 'billing_city', 'string' );
        $this->addColumn( 'billing_discount', 'int' );
        $this->addColumn( 'billing_discount_code', 'string' );
        $this->addColumn( 'billing_email', 'string' );
        $this->addColumn( 'billing_name', 'string' );
        $this->addColumn( 'billing_name_on_card', 'string' );
        $this->addColumn( 'billing_phone', 'string' );
        $this->addColumn( 'billing_postalcode', 'string' );
        $this->addColumn( 'billing_province', 'string' );
        $this->addColumn( 'billing_subtotal', 'int' );
        $this->addColumn( 'billing_tax', 'int' );
        $this->addColumn( 'billing_total', 'int' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'error', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'payment_gateway', 'string' );
        $this->addColumn( 'shipped', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

