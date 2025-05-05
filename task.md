# Transactional outbox with orders-center

В проде у нас есть проект **orders-center** - это сервис для работы с заказами, являющийся единственным источником истины по статусам и информации о заказе.

В процессе развития компании мы плавно переходим с 1С на веб микросервисы.
Когда в **orders-center** прилетает запрос на сохранение заказа из другого микросервиса, мы должны синхронизировать этот заказ с 1С (отправить запрос на сохранение по REST).

# Доменные модели
## Order

    type Order struct {  
      ID string    
      Type string    
      Status string     
      City string    
      Subdivision string    
      Price float64   
      Platform string      
      GeneralID uuid.UUID 
      OrderNumber string    
      Executor string     
      CreatedAt time.Time 
      UpdatedAt time.Time
    }


## History

    type History struct {  
      Type string    `json:"type"`  
      TypeId int       `json:"type_id"`  
      OldValue []byte    `json:"old_value"`  
      Value []byte    `json:"value"`  
      Date time.Time `json:"date"`  
      UserID string    `json:"user_id"`  
      OrderID string    `json:"order_id"`  
    }

## Order Item

    OrderItem struct {  
      ProductID string   
      ExternalID string    
      Status string   
      BasePrice float64  
      Price float64  
      EarnedBonuses float64   
      SpentBonuses float64  
      Gift bool     
      OwnerID string    
      DeliveryID string   
      ShopAssistant string    
      Warehouse string  
      OrderId uuid.UUID 
    }


## Payment

    type OrderPayment struct {  
      ID uuid.UUID      
      OrderID uuid.UUID       
      Type PaymentType     
      Sum float64          
      Payed bool            
      Info string         
      CreditData *CreditData      
      ContractNumber string         
      CardPaymentData *CardPaymentData 
      ExternalID string           
    }

    type PaymentType string  
      
    const (  
      PaymentTypeCashAtShop PaymentType = "cash_at_shop"  
      PaymentTypeCashToCourier PaymentType = "cash_to_courier"  
      PaymentTypeCard PaymentType = "card"  
      PaymentTypeCardOnline PaymentType = "card_online"  
      PaymentTypeCredit PaymentType = "credit"  
      PaymentTypeBonuses PaymentType = "bonuses"  
      PaymentTypeCashless PaymentType = "cashless"  
      PaymentTypePrepayment PaymentType = "prepayment"  
    )

    type CreditData struct {  
      Bank string  
      Type string  
      NumberOfMonths int16   
      PaySumPerMonth float64 
      BrokerID int32
      IIN string  
    }
    
    type CardPaymentData struct {  
      Provider string 
      TransactionId string 
    }


# Задачи

1.  Создать общую структуру для хранения всего заказа вместе, назвать ее **OrderFull**.
2. Написать мок сервис **1С** (отдельным проектом), который принимает запрос по rest и сохраняет у себя данные заказа. (можно без бд, просто хранить в оперативке)
3. Написать и реализовать сервис **order_eno_1c** внутри **orders-center** который сохраняет таску в бд, может получать необработанные таски и процессить их.
4. Пусть **cron** будет реализован через **worker pool**
5. Пусть будут соблюдены условия конкурентности, например если несколько экземпляров сервиса одновременно начинают читать таски из бд.
6. Пусть **order_eno_1c** будет инжектить **cron** в себя и ничего не знать о реализации **worker_pool**
7. Покрыть код unit тестами