package com.loginapi.loginapi.security;

import io.jsonwebtoken.*;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class JwtUtil {
    //Lấy trong .properties
    @Value("${jwt.secret}")
    private String secret;
    // Tạo 1 token
    public String generateToken(Integer userId) {
        return Jwts.builder()
                .claim("userId", userId)
                .expiration(new Date(System.currentTimeMillis() + 86400000))
                .signWith(Keys.hmacShaKeyFor(secret.getBytes()))
                .compact();
    }

    public Integer parseToken(String token) {
        Claims claims = Jwts.parser()
                .verifyWith(Keys.hmacShaKeyFor(secret.getBytes()))
                .build()
                .parseSignedClaims(token)
                .getPayload();
        return claims.get("userId",Integer.class);
    }

}