




/**
 *  @module : adc1
 *
 *  Initializes and samples from the ADC1 peripheral across all channels
 *  The inputs to this API is a gpio enumeration which maps directly and uniquely to an ADC1 channel
 */

/**
 *  Initializes an ADC1 channel
 *  @param gpio : The GPIO to select as an ADC mux and map to an ADC channel
 *  @returns    : True for successful, flase for not successful
 */
bool adc1_initialize(const gpio_E gpio);

/**
 *  Samples from the ADC register
 *  @param gpio : The GPIO to map to the ADC channel to read from
 *  @returns    : ADC value in volts, negative represents an invalid sample
 */
int32_t acd1_sample(const gpio_E gpio);

/**
 *  Checks if a certain GPIO is initializd for ADC1
 *  @param gpio : The GPIO to check
 *  @returns    : True for initialized, false for not
 */
bool adc1_is_initialized(const gpio_E gpio);

/// Stores a map of bools to keep track of which ADC 1 channels has been initialized
static bool adc_init_map[ADC1_CHANNEL_MAX] = { 0 };

/**
 *  Converts a GPIO enumeration to an ADC 1 channel enumeration
 *  Always a valid enumeration but user must check if it is MAX or not
 *  @param gpio : The input to convert
 *  @returns    : An ADC 1 channel
 */
static adc1_channel_t adc_gpio_to_channel(const gpio_E gpio)
{
    switch (gpio_get_pin_number(gpio))
    {
        case GPIO_NUM_36: return ADC1_CHANNEL_0;
        case GPIO_NUM_37: return ADC1_CHANNEL_1;
        case GPIO_NUM_38: return ADC1_CHANNEL_2;
        case GPIO_NUM_39: return ADC1_CHANNEL_3;
        case GPIO_NUM_32: return ADC1_CHANNEL_4;
        case GPIO_NUM_33: return ADC1_CHANNEL_5;
        case GPIO_NUM_34: return ADC1_CHANNEL_6;
        case GPIO_NUM_35: return ADC1_CHANNEL_7;
        default:          return ADC1_CHANNEL_MAX;
    }
}

bool adc1_initialize(const gpio_E gpio)
{
    const adc1_channel_t channel = adc_gpio_to_channel(gpio);
    bool success = false;

    // If not already initialized
    if (adc_init_map[channel])
    {
        success = true;
    }
    else if (channel < ADC1_CHANNEL_MAX)
    {
        // Enable GPIO
        const adc_unit_t adc_unit = ADC_UNIT_1;
        adc_gpio_init(adc_unit, channel);

        // Set ADC resolution
        const adc_bits_width_t adc_resolution = ADC_WIDTH_12Bit;
        ESP_ERROR_CHECK(adc1_config_width(adc_resolution));

        // Enable ADC and set attenuation
        const adc_atten_t full_scale_attenuation = ADC_ATTEN_DB_6;
        ESP_ERROR_CHECK(adc1_config_channel_atten(channel, full_scale_attenuation));
        adc_init_map[channel] = true;

        success = true;
    }
    else
    {
        ESP_LOGE("adc1_initialize", "Incorrect channel %d", channel);
    }

    return success;
}

int32_t acd1_sample(const gpio_E gpio)
{
    const adc1_channel_t channel = adc_gpio_to_channel(gpio);
    int32_t adc_voltage_v = -1;

    if (channel < ADC1_CHANNEL_MAX && adc_init_map[channel])
    {
        adc_voltage_v = adc1_get_raw(channel);
    }

    return adc_voltage_v;
}

bool adc1_is_initialized(const gpio_E gpio)
{
    return adc_init_map[adc_gpio_to_channel(gpio)];
}
