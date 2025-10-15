<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait WebsocketsStatisticsEntryColumnsTraits
{
    public function apiMessageCount()
    {
        $this->addColumn( 'api_message_count' );
        return $this;
    }

    public function appId()
    {
        $this->addColumn( 'app_id' );
        return $this;
    }

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

    public function peakConnectionCount()
    {
        $this->addColumn( 'peak_connection_count' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function websocketMessageCount()
    {
        $this->addColumn( 'websocket_message_count' );
        return $this;
    }

}

