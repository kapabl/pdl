<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait SessionColumnsTraits
{
    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function ipAddress()
    {
        $this->addColumn( 'ip_address' );
        return $this;
    }

    public function lastActivity()
    {
        $this->addColumn( 'last_activity' );
        return $this;
    }

    public function payload()
    {
        $this->addColumn( 'payload' );
        return $this;
    }

    public function userAgent()
    {
        $this->addColumn( 'user_agent' );
        return $this;
    }

    public function userId()
    {
        $this->addColumn( 'user_id' );
        return $this;
    }

}

