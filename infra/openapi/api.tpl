consumes:
- application/json
definitions:
  Address:
    properties:
      city:
        type: string
        x-omitempty: true
      country:
        type: string
        x-omitempty: true
      country_code:
        type: string
        x-omitempty: true
      extended_address:
        type: string
        x-omitempty: true
      po_box:
        type: string
        x-omitempty: true
      postal_code:
        type: string
        x-omitempty: true
      region:
        type: string
        x-omitempty: true
      street_address:
        type: string
        x-omitempty: true
      type:
        type: string
        x-omitempty: true
    type: object
  Birthday:
    properties:
      date:
        type: string
        x-omitempty: true
      text:
        type: string
        x-omitempty: true
    type: object
  Contact:
    properties:
      addresses:
        items:
          $ref: '#/definitions/Address'
        type: array
      birthdays:
        items:
          $ref: '#/definitions/Birthday'
        type: array
      category:
        $ref: '#/definitions/ContactCategory'
      email_addresses:
        items:
          $ref: '#/definitions/EmailAddress'
        type: array
      genders:
        items:
          $ref: '#/definitions/Gender'
        type: array
      id:
        type: string
      last_contact:
        format: date-time
        type: string
      names:
        items:
          $ref: '#/definitions/UserNames'
        type: array
      next_contact:
        format: date-time
        type: string
      nicknames:
        items:
          $ref: '#/definitions/Nickname'
        type: array
      occupations:
        items:
          $ref: '#/definitions/Occupation'
        type: array
      organizations:
        items:
          $ref: '#/definitions/Organization'
        type: array
      phone_numbers:
        items:
          $ref: '#/definitions/PhoneNumber'
        type: array
      photos:
        items:
          $ref: '#/definitions/Photo'
        type: array
      relations:
        items:
          $ref: '#/definitions/Relation'
        type: array
      score:
        type: integer
      urls:
        items:
          $ref: '#/definitions/Url'
        type: array
    title: User
    type: object
  ContactCategory:
    enum:
    - A
    - B
    - C
    - D
    type: string
  ContactSource:
    properties:
      created_at:
        format: date-time
        type: string
      email:
        type: string
      id:
        type: string
      source:
        type: string
      updated_at:
        format: date-time
        type: string
      user_id:
        type: string
    title: contact-source
    type: object
  CreateContactDto:
    properties:
      addresses:
        items:
          $ref: '#/definitions/Address'
        type: array
      birthdays:
        items:
          $ref: '#/definitions/Birthday'
        type: array
      display_name:
        type: string
      email_addresses:
        items:
          $ref: '#/definitions/EmailAddress'
        type: array
      genders:
        items:
          $ref: '#/definitions/Gender'
        type: array
      last_contact:
        format: date-time
        type: string
      names:
        items:
          $ref: '#/definitions/UserNames'
        type: array
      next_contact:
        format: date-time
        type: string
      nicknames:
        items:
          $ref: '#/definitions/Nickname'
        type: array
      occupations:
        items:
          $ref: '#/definitions/Occupation'
        type: array
      organizations:
        items:
          $ref: '#/definitions/Organization'
        type: array
      phone_numbers:
        items:
          $ref: '#/definitions/PhoneNumber'
        type: array
      photos:
        items:
          $ref: '#/definitions/Photo'
        type: array
      relations:
        items:
          $ref: '#/definitions/Relation'
        type: array
      score:
        type: integer
      urls:
        items:
          $ref: '#/definitions/Url'
        type: array
    title: CreateContactDto
    type: object
  EmailAddress:
    properties:
      display_name:
        type: string
        x-omitempty: true
      type:
        type: string
        x-omitempty: true
      value:
        type: string
        x-omitempty: true
    type: object
  ErrorResponse:
    properties:
      description:
        type: string
      error:
        type: string
    title: ErrorResponse
    type: object
  Gender:
    properties:
      address_me_as:
        type: string
        x-omitempty: true
      value:
        type: string
        x-omitempty: true
    type: object
  InitResponse:
    properties:
      url:
        type: string
    title: init response
    type: object
  LinkMatch:
    properties:
      display_name:
        type: string
      unified_id:
        type: string
    title: LinkMatch
    type: object
  LinkSuggestion:
    properties:
      id:
        type: string
      key:
        type: string
      matches:
        items:
          $ref: '#/definitions/LinkMatch'
        type: array
      value:
        type: string
    title: LinkSuggestion
    type: object
  Message:
    properties:
      message:
        type: string
    title: message
    type: object
  Nickname:
    properties:
      value:
        type: string
        x-omitempty: true
    type: object
  Note:
    properties:
      created_at:
        format: date-time
        type: string
      id:
        type: string
      is_updated:
        type: boolean
      note:
        type: string
    title: note
    type: object
  Occupation:
    properties:
      value:
        type: string
        x-omitempty: true
    type: object
  Organization:
    properties:
      department:
        type: string
        x-omitempty: true
      domain:
        type: string
      end_date:
        type: string
        x-omitempty: true
      is_current:
        type: boolean
        x-omitempty: true
      job_description:
        type: string
        x-omitempty: true
      location:
        type: string
        x-omitempty: true
      name:
        type: string
        x-omitempty: true
      phonetic_name:
        type: string
        x-omitempty: true
      start_date:
        type: string
        x-omitempty: true
      symbol:
        type: string
        x-omitempty: true
      title:
        type: string
        x-omitempty: true
      type:
        type: string
        x-omitempty: true
    type: object
  PhoneNumber:
    properties:
      type:
        type: string
        x-omitempty: true
      value:
        type: string
        x-omitempty: true
    type: object
  Photo:
    properties:
      default:
        type: boolean
        x-omitempty: true
      url:
        type: string
        x-omitempty: true
    type: object
  Quota:
    properties:
      total_category_assigned:
        type: integer
      total_contacts:
        type: integer
    title: quota
    type: object
  Relation:
    properties:
      person:
        type: string
        x-omitempty: true
      type:
        type: string
        x-omitempty: true
    type: object
  SearchContactDto:
    properties:
      filters:
        items:
          $ref: '#/definitions/SearchFilter'
        type: array
      page:
        type: number
      per_page:
        type: number
      query:
        type: string
      sort:
        items:
          $ref: '#/definitions/SearchSort'
        type: array
    title: SearchContactDto
    type: object
  SearchFilter:
    properties:
      field:
        enum:
        - category
        - next_contact
        - last_contact
        - score
        - birthday
        - gender
        - tag
        type: string
      operator:
        type: string
      value:
        type: string
    type: object
  SearchSort:
    properties:
      field:
        type: string
      order:
        enum:
        - asc
        - desc
        type: string
    type: object
  Tag:
    properties:
      created_at:
        format: date-time
        type: string
      id:
        type: string
      tag_name:
        type: string
    title: tag
    type: object
  Unified:
    properties:
      addresses:
        items:
          $ref: '#/definitions/Address'
        type: array
      birthdays:
        items:
          $ref: '#/definitions/Birthday'
        type: array
      category:
        enum:
        - A
        - B
        - C
        - D
        type: string
      display_name:
        type: string
      email_addresses:
        items:
          $ref: '#/definitions/EmailAddress'
        type: array
      genders:
        items:
          $ref: '#/definitions/Gender'
        type: array
      id:
        type: string
      last_contact:
        format: date-time
        type: string
      names:
        items:
          $ref: '#/definitions/UserNames'
        type: array
      next_contact:
        format: date-time
        type: string
      nicknames:
        items:
          $ref: '#/definitions/Nickname'
        type: array
      occupations:
        items:
          $ref: '#/definitions/Occupation'
        type: array
      organizations:
        items:
          $ref: '#/definitions/Organization'
        type: array
      phone_numbers:
        items:
          $ref: '#/definitions/PhoneNumber'
        type: array
      photos:
        items:
          $ref: '#/definitions/Photo'
        type: array
      relations:
        items:
          $ref: '#/definitions/Relation'
        type: array
      score:
        type: integer
      urls:
        items:
          $ref: '#/definitions/Url'
        type: array
    title: Unified Contact
    type: object
  UpdateCategoryDto:
    properties:
      category:
        $ref: '#/definitions/ContactCategory'
    title: UpdateCategoryDto
    type: object
  UpdateUnifiedDto:
    properties:
      addresses:
        items:
          $ref: '#/definitions/Address'
        type: array
      birthdays:
        items:
          $ref: '#/definitions/Birthday'
        type: array
      display_name:
        type: string
      email_addresses:
        items:
          $ref: '#/definitions/EmailAddress'
        type: array
      genders:
        items:
          $ref: '#/definitions/Gender'
        type: array
      last_contact:
        format: date-time
        type: string
      names:
        items:
          $ref: '#/definitions/UserNames'
        type: array
      next_contact:
        format: date-time
        type: string
      nicknames:
        items:
          $ref: '#/definitions/Nickname'
        type: array
      occupations:
        items:
          $ref: '#/definitions/Occupation'
        type: array
      organizations:
        items:
          $ref: '#/definitions/Organization'
        type: array
      phone_numbers:
        items:
          $ref: '#/definitions/PhoneNumber'
        type: array
      photos:
        items:
          $ref: '#/definitions/Photo'
        type: array
      relations:
        items:
          $ref: '#/definitions/Relation'
        type: array
      score:
        type: integer
      urls:
        items:
          $ref: '#/definitions/Url'
        type: array
    title: UpdateUnifiedDto
    type: object
  Url:
    properties:
      type:
        type: string
        x-omitempty: true
      value:
        type: string
        x-omitempty: true
    type: object
  User:
    properties:
      addresses:
        items:
          $ref: '#/definitions/Address'
        type: array
      birthdays:
        items:
          $ref: '#/definitions/Birthday'
        type: array
      email_addresses:
        items:
          $ref: '#/definitions/EmailAddress'
        type: array
      genders:
        items:
          $ref: '#/definitions/Gender'
        type: array
      id:
        type: string
      names:
        items:
          $ref: '#/definitions/UserNames'
        type: array
      nicknames:
        items:
          $ref: '#/definitions/Nickname'
        type: array
      occupations:
        items:
          $ref: '#/definitions/Occupation'
        type: array
      organizations:
        items:
          $ref: '#/definitions/Organization'
        type: array
      phone_numbers:
        items:
          $ref: '#/definitions/PhoneNumber'
        type: array
      photos:
        items:
          $ref: '#/definitions/Photo'
        type: array
      quota:
        $ref: '#/definitions/Quota'
      relations:
        items:
          $ref: '#/definitions/Relation'
        type: array
      urls:
        items:
          $ref: '#/definitions/Url'
        type: array
    title: User
    type: object
  UserNames:
    properties:
      display_name:
        type: string
        x-omitempty: true
      display_name_last_first:
        type: string
        x-omitempty: true
      family_name:
        type: string
        x-omitempty: true
      given_name:
        type: string
        x-omitempty: true
      honorific_prefix:
        type: string
        x-omitempty: true
      honorific_suffix:
        type: string
        x-omitempty: true
      middle_name:
        type: string
        x-omitempty: true
      phonetic_family_name:
        type: string
        x-omitempty: true
      phonetic_full_name:
        type: string
        x-omitempty: true
      phonetic_given_name:
        type: string
        x-omitempty: true
      phonetic_honorific_prefix:
        type: string
        x-omitempty: true
      phonetic_honorific_suffix:
        type: string
        x-omitempty: true
      phonetic_middle_name:
        type: string
        x-omitempty: true
      unstructured_name:
        type: string
        x-omitempty: true
    type: object
info:
  description: contact karma service
  title: Contact Karma Service
  version: 1.0.0
paths:
  /contacts/sources:
    get:
      operationId: get-contact-sources
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/ContactSource'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get user's contact source list
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-contact-sources
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /contacts/sources/{source_id}:
    delete:
      operationId: delete-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: source id
        in: path
        name: source_id
        required: true
        type: string
      responses:
        "200":
          description: Deleted
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: delete contact source by id
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-delete-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: source id
        in: path
        name: source_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /contacts/sources/google/init:
    get:
      operationId: init-google-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: Init Response
          schema:
            $ref: '#/definitions/InitResponse'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: returns redirect url
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-init-google-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /contacts/sources/google/link:
    options:
      operationId: cors-link-google-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: link-google-contact-source
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: body
        name: body
        schema:
          properties:
            auth_code:
              type: string
          type: object
      responses:
        "200":
          description: Linked
          headers:
            Location:
              type: string
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "403":
          description: quota limit reached
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: redirects user to google consent page
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /contacts/upload-csv:
    options:
      operationId: cors-upload-contacts-csv
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      consumes:
      - multipart/form-data
      operationId: upload-contacts-csv
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: The uploaded file data
        format: binary
        in: formData
        name: file
        required: true
        type: string
      responses:
        "200":
          description: Uploaded
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: upload contacts csv
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /pending:
    get:
      operationId: get-pending-follow-ups
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - format: date-time
        in: query
        name: last_document_next_contact
        type: string
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/Unified'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get pending follow ups
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-pending-follow-ups
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /recent:
    get:
      operationId: get-recent-contacts
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: query
        name: max_days
        type: integer
      - in: query
        name: limit
        type: integer
      - format: date-time
        in: query
        name: last_document_last_contact
        type: string
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/Unified'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get pending follow ups
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-recent-contacts
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /search:
    options:
      operationId: cors-search-user-contact
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: search-user-contact
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: body
        name: body
        required: true
        schema:
          $ref: '#/definitions/SearchContactDto'
      responses:
        "200":
          description: Created
          schema:
            items:
              $ref: '#/definitions/Unified'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: search user's contact
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /suggestions:
    get:
      operationId: get-link-suggestions
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/LinkSuggestion'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get list of contact link suggestions to remove duplicates
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-link-suggestions
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /suggestions/{suggestion_id}/apply:
    options:
      operationId: cors-apply-link-suggestion
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: suggestion id
        in: path
        name: suggestion_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: apply-link-suggestion
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: suggestion id
        in: path
        name: suggestion_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          properties:
            unified_ids:
              items:
                type: string
              type: array
          type: object
      responses:
        "200":
          description: Applied
          headers:
            Location:
              type: string
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: redirects user to google consent page
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified:
    get:
      operationId: get-unified-contacts
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/Unified'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get unified list of contacts
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-unified-contacts
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: create-user-contact
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/CreateContactDto'
      responses:
        "200":
          description: Created
          schema:
            $ref: '#/definitions/Unified'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: create user's contact
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}:
    delete:
      operationId: delete-user-contact
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: Deleted
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: delete user's contact by id
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    get:
      operationId: get-user-contact-by-id
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            $ref: '#/definitions/Unified'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get user's contact by id
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-user-contact-by-id
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    patch:
      operationId: update-user-contact
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/UpdateUnifiedDto'
      responses:
        "200":
          description: Updated
          schema:
            $ref: '#/definitions/Unified'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: update user's contact by id
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}/category:
    options:
      operationId: cors-update-contact-category
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    patch:
      operationId: update-contact-category
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/UpdateCategoryDto'
      responses:
        "200":
          description: Updated
          schema:
            $ref: '#/definitions/Unified'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: update category
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}/notes:
    get:
      operationId: get-contact-notes
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/Note'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get contact notes
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-contact-notes
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: post-contact-note
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/Note'
      responses:
        "200":
          description: Created
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: post contact note
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}/notes/{note_id}:
    delete:
      operationId: delete-contact-note
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: note id
        in: path
        name: note_id
        required: true
        type: string
      responses:
        "200":
          description: Deleted
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: delete contact note
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-patch-contact-note
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: note id
        in: path
        name: note_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    patch:
      operationId: patch-contact-note
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: note id
        in: path
        name: note_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/Note'
      responses:
        "200":
          description: Updated
          schema:
            $ref: '#/definitions/Note'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: patch contact note
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}/tags:
    get:
      operationId: get-contact-tags
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: query
        name: limit
        type: integer
      - in: query
        name: last_document_id
        type: string
      responses:
        "200":
          description: Retrieved
          schema:
            items:
              $ref: '#/definitions/Tag'
            type: array
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get contact tags
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-contact-tags
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: post-contact-tag
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/Tag'
      responses:
        "200":
          description: Created
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: post contact tag
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /unified/{unified_id}/tags/{tag_id}:
    delete:
      operationId: delete-contact-tag
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: tag id
        in: path
        name: tag_id
        required: true
        type: string
      responses:
        "200":
          description: Deleted
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: delete contact tag
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-patch-contact-tag
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: tag id
        in: path
        name: tag_id
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    patch:
      operationId: patch-contact-tag
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - description: contact id
        in: path
        name: unified_id
        required: true
        type: string
      - description: tag id
        in: path
        name: tag_id
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/Tag'
      responses:
        "200":
          description: Updated
          schema:
            $ref: '#/definitions/Tag'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "404":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: patch contact tag
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /users:
    options:
      operationId: cors-create-user
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    patch:
      operationId: update-user
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/User'
      responses:
        "200":
          description: Updated
          schema:
            $ref: '#/definitions/User'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: update user
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    post:
      operationId: create-user
      parameters:
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      - in: body
        name: body
        schema:
          $ref: '#/definitions/User'
      responses:
        "200":
          description: Created
          schema:
            $ref: '#/definitions/User'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: create user
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
  /users/{user_id}:
    delete:
      operationId: delete-user
      parameters:
      - description: user id
        in: path
        name: user_id
        required: true
        type: string
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: Deleted
          schema:
            $ref: '#/definitions/Message'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: delete user
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    get:
      operationId: get-user
      parameters:
      - description: user id
        in: path
        name: user_id
        required: true
        type: string
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: retrieved
          schema:
            $ref: '#/definitions/User'
        "400":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "401":
          description: Error
          schema:
            $ref: '#/definitions/ErrorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/ErrorResponse'
      security:
      - firebase: []
      summary: get user
      x-google-backend:
        address: ${contacts_url}
        path_translation: APPEND_PATH_TO_ADDRESS
    options:
      operationId: cors-get-user
      parameters:
      - description: user id
        in: path
        name: user_id
        required: true
        type: string
      - in: header
        name: X-Apigateway-Api-Userinfo
        required: true
        type: string
      responses:
        "200":
          description: A successful response
      x-google-backend:
        address: ${options_url}
        path_translation: APPEND_PATH_TO_ADDRESS
produces:
- application/json
schemes:
- http
- https
securityDefinitions:
  firebase:
    authorizationUrl: ""
    flow: implicit
    scopes: {}
    type: oauth2
    x-google-audiences: ${project_name}
    x-google-issuer: https://securetoken.google.com/${project_name}
    x-google-jwks_uri: https://www.googleapis.com/service_accounts/v1/metadata/x509/securetoken@system.gserviceaccount.com
swagger: "2.0"
