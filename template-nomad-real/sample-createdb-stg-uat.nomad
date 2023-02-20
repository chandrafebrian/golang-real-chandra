#
#sample real create database staging-uat-pre-prduction lewat hcl static job

job "postgres-service-notif--preprod" {
  datacenters = ["staging"]
  
  constraint {
    attribute = "${meta.node_type}"
    value     = "ondemand"
  }

  group "postgresql" {
    count = 1

    volume "postgres-service-notif--preprod" {
      type      = "csi"
      read_only = false
      source    = "postgres-service-notif--preprod"
      access_mode     = "single-node-writer"
      attachment_mode = "file-system"
    }

    network { 
      port "postgres-service-notif--preprod" {
        to = 5432
        static = 7890
      }
      dns {
        servers = ["172.31.0.191"]
      }
    }

    task "postgresql" {
      driver = "docker"

      volume_mount {
        volume      = "postgres-service-notif--preprod"
        destination = "/var/lib/postgresql/data"
        read_only   = false
      }

      env {
        POSTGRES_DB = "postgres"
        POSTGRES_USER = "postgres"
        POSTGRES_PASSWORD = "YITt5zLLy4"
        PGDATA = "/var/lib/postgresql/data/pgdata"
      }

      config {
        image = "postgres:15"
        ports = ["postgres-service-notif--preprod"]
      }

      resources {
        cpu    = 1024
        memory = 1024
      }
    }

    service {
      name = "postgres-service-notif--preprod"
      port = "postgres-service-notif--preprod"

      check {
        port        = "postgres-service-notif--preprod"
        type        = "tcp"
        interval    = "15s"
        timeout     = "10s"
      }
    }
  }
}