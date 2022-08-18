package com.fabrictemplate.template.dto;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;


@Data
@AllArgsConstructor
@NoArgsConstructor
public class Person {

    @JsonInclude(JsonInclude.Include.NON_NULL)
    private String firstName;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    private String lastName;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    private int age;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    private String nationalId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    private List<SocialMediaAccounts> socialMediaAccounts;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    private boolean isMarried;
}
