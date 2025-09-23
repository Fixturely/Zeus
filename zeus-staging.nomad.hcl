job "zeus-staging" {
  datacenters = ["dc1"]
  type = "service"

  group "zeus-staging" {
    count = 1

    network {
      port "http" {
        static = 8081
        to     = 8080
      }
    }

    service {
      name = "zeus-staging"
      port = "http"
      
      check {
        name     = "health"
        type     = "http"
        path     = "/health"
        interval = "10s"
        timeout  = "2s"
      }
    }

    task "zeus-staging" {
      driver = "exec"

      config {
        command = "/local/zeus"
        args    = []
      }

      resources {
        cpu    = 250
        memory = 256
      }

      env {
        ENV = "staging"
        PORT = "8080"
      }

      artifact {
        source = "file://deploy/staging/zeus-staging-release.tar.gz"
        destination = "/local"
        mode = "file"
      }

      logs {
        max_files     = 3
        max_file_size = 5
      }
    }
  }

  update {
    max_parallel     = 1
    health_check     = "checks"
    min_healthy_time = "10s"
    healthy_deadline = "3m"
    auto_revert      = true
    canary           = 0
  }
}
