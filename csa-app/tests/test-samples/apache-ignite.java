package com.bsfit.data.service.model;

import java.io.Serializable;
import org.apache.ignite.cache.query.annotations.QuerySqlField;


/**
 * 
 * @author Jet
 *
 */
public class JetAddress implements Serializable {
    
    /**
     * 
     */
    private static final long serialVersionUID = 7538875097713370177L;

    @QuerySqlField
    public long id;
    
    @QuerySqlField(index = true)
    public String province;

    @QuerySqlField(index = true)
    public String city;

    @QuerySqlField(index = true)
    public String county;

    @QuerySqlField(index = true)
    public String township;

    @QuerySqlField
    public String restaddr;
    
    
    /**
     * Default constructor.
     */
    public JetAddress() {
        // No-op.
    }


    public long getId() {
        return id;
    }


    public void setId(long id) {
        this.id = id;
    }


    public String getProvince() {
        return province;
    }


    public void setProvince(String province) {
        this.province = province;
    }


    public String getTownship() {
        return township;
    }


    public void setTownship(String township) {
        this.township = township;
    }


    public String getRestaddr() {
        return restaddr;
    }


    public void setRestaddr(String restaddr) {
        this.restaddr = restaddr;
    }
    
    
}
