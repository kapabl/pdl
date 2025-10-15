<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait ProductCustomizationRelationColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function parentId()
    {
        $this->addColumn( 'parent_id' );
        return $this;
    }

    public function position()
    {
        $this->addColumn( 'position' );
        return $this;
    }

    public function productCustomizationId()
    {
        $this->addColumn( 'product_customization_id' );
        return $this;
    }

    public function productId()
    {
        $this->addColumn( 'product_id' );
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

