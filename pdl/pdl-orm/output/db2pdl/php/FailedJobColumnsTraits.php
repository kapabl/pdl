<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait FailedJobColumnsTraits
{
    public function connection()
    {
        $this->addColumn( 'connection' );
        return $this;
    }

    public function exception()
    {
        $this->addColumn( 'exception' );
        return $this;
    }

    public function failedAt()
    {
        $this->addColumn( 'failed_at' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function payload()
    {
        $this->addColumn( 'payload' );
        return $this;
    }

    public function queue()
    {
        $this->addColumn( 'queue' );
        return $this;
    }

    public function uuid()
    {
        $this->addColumn( 'uuid' );
        return $this;
    }

}

