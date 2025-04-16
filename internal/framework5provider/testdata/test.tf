resource "framework_identity" "test" {
  name = "john"
}

import {
  to = framework_identity.test
  identity = {
    id = "id-123"
  }
}
