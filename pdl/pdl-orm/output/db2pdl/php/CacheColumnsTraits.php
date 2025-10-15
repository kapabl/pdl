<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait CacheColumnsTraits
{
    public function expiration()
    {
        $this->addColumn( 'expiration' );
        return $this;
    }

    public function key()
    {
        $this->addColumn( 'key' );
        return $this;
    }

    public function value()
    {
        $this->addColumn( 'value' );
        return $this;
    }

}

