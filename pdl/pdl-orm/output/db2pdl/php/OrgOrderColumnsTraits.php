<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait OrgOrderColumnsTraits
{
    public function billingAddress()
    {
        $this->addColumn( 'billing_address' );
        return $this;
    }

    public function billingCity()
    {
        $this->addColumn( 'billing_city' );
        return $this;
    }

    public function billingDiscount()
    {
        $this->addColumn( 'billing_discount' );
        return $this;
    }

    public function billingDiscountCode()
    {
        $this->addColumn( 'billing_discount_code' );
        return $this;
    }

    public function billingEmail()
    {
        $this->addColumn( 'billing_email' );
        return $this;
    }

    public function billingName()
    {
        $this->addColumn( 'billing_name' );
        return $this;
    }

    public function billingNameOnCard()
    {
        $this->addColumn( 'billing_name_on_card' );
        return $this;
    }

    public function billingPhone()
    {
        $this->addColumn( 'billing_phone' );
        return $this;
    }

    public function billingPostalcode()
    {
        $this->addColumn( 'billing_postalcode' );
        return $this;
    }

    public function billingProvince()
    {
        $this->addColumn( 'billing_province' );
        return $this;
    }

    public function billingSubtotal()
    {
        $this->addColumn( 'billing_subtotal' );
        return $this;
    }

    public function billingTax()
    {
        $this->addColumn( 'billing_tax' );
        return $this;
    }

    public function billingTotal()
    {
        $this->addColumn( 'billing_total' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function error()
    {
        $this->addColumn( 'error' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function paymentGateway()
    {
        $this->addColumn( 'payment_gateway' );
        return $this;
    }

    public function shipped()
    {
        $this->addColumn( 'shipped' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function userId()
    {
        $this->addColumn( 'user_id' );
        return $this;
    }

}

