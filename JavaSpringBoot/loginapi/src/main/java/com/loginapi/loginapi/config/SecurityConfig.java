package com.loginapi.loginapi.config;

import com.loginapi.loginapi.security.JwtFilter;
import lombok.RequiredArgsConstructor;
import org.springframework.context.annotation.*;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.web.*;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

@Configuration
@RequiredArgsConstructor
public class SecurityConfig {

    private final JwtFilter jwt;
    @Bean
    SecurityFilterChain filter(HttpSecurity http)throws Exception {
        http.csrf(c -> c.disable())
                .authorizeHttpRequests(a -> a.requestMatchers("/api/auth/login")
                        .permitAll()
                        .anyRequest()
                        .permitAll()
                )
                //  Khi có request gửi đến, chạy qua JwtFilter trước
                .addFilterBefore(jwt, UsernamePasswordAuthenticationFilter.class);
        return http.build();

    }

}