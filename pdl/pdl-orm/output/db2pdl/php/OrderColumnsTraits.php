<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait OrderColumnsTraits
{
    public function courierUuid()
    {
        $this->addColumn( 'courier_uuid' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function deliveryTaxAmount()
    {
        $this->addColumn( 'delivery_tax_amount' );
        return $this;
    }

    public function discountAmount()
    {
        $this->addColumn( 'discount_amount' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function model()
    {
        $this->addColumn( 'model' );
        return $this;
    }

    public function notes()
    {
        $this->addColumn( 'notes' );
        return $this;
    }

    public function orderId()
    {
        $this->addColumn( 'order_id' );
        return $this;
    }

    public function sellerTaxAmount()
    {
        $this->addColumn( 'seller_tax_amount' );
        return $this;
    }

    public function sellerTotal()
    {
        $this->addColumn( 'seller_total' );
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

    public function subTotal()
    {
        $this->addColumn( 'sub_total' );
        return $this;
    }

    public function taxAmount()
    {
        $this->addColumn( 'tax_amount' );
        return $this;
    }

    public function total()
    {
        $this->addColumn( 'total' );
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

