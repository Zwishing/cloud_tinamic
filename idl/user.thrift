namespace go base.user

include "base.thrift"

struct User {
    1: string user_id,
    2: string name,
    3: bool state
    4: string avatar,
    5: string phone_number,
    6: string salt,
    7: string password,
}

enum UserCategory {
    USERNAME = 1,
    EMAIL = 2,
    PHONE = 3
}
struct Account{
    1: i64 user_id,
    2: string username,
    3: UserCategory user_category
}

enum RoleCode{
    SUPERADMIN = 1,
    ADMIN = 2,
    USER = 3,
    GUEST = 4,
}

struct Role {
     1: i64 id,
     2: i64 parent_id,
     3: string name,
     4: RoleCode role_code,
}

struct UserRole {
    1: string user_id
    2: RoleCode role_code
}

struct RegisterRequest {
    1: string username,
    2: string password,
    3: UserCategory user_category,
    4: RoleCode role_code,
}

struct RegisterResponse {
    1: base.BaseResp base,
    2: string user_id,
    3: string token,
}

struct LoginRequest {
    1: string username,
    2: string password,
    3: UserCategory user_category,
}

struct LoginResponse {
    1: base.BaseResp base,
    2: User user,
    3: string token,
}

struct InfoRequest {
    1: i64 user_id,
    2: string token,
}

struct InfoResponse {
    1: base.BaseResp base,
    2: User user,
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    InfoResponse Info(1: InfoRequest req),
}