<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class OrgOrderOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'billing_address' );
        $this->addField( 'billing_city' );
        $this->addField( 'billing_discount' );
        $this->addField( 'billing_discount_code' );
        $this->addField( 'billing_email' );
        $this->addField( 'billing_name' );
        $this->addField( 'billing_name_on_card' );
        $this->addField( 'billing_phone' );
        $this->addField( 'billing_postalcode' );
        $this->addField( 'billing_province' );
        $this->addField( 'billing_subtotal' );
        $this->addField( 'billing_tax' );
        $this->addField( 'billing_total' );
        $this->addField( 'created_at' );
        $this->addField( 'error' );
        $this->addField( 'id' );
        $this->addField( 'payment_gateway' );
        $this->addField( 'shipped' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
    }
}

