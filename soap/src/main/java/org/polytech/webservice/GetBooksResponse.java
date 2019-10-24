
package org.polytech.webservice;

import javax.xml.bind.annotation.XmlAccessType;
import javax.xml.bind.annotation.XmlAccessorType;
import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;
import javax.xml.bind.annotation.XmlType;


/**
 * <p>Classe Java pour anonymous complex type.
 * 
 * <p>Le fragment de schéma suivant indique le contenu attendu figurant dans cette classe.
 * 
 * <pre>
 * &lt;complexType>
 *   &lt;complexContent>
 *     &lt;restriction base="{http://www.w3.org/2001/XMLSchema}anyType">
 *       &lt;sequence>
 *         &lt;element name="books" type="{http://spring.io/guides/gs-producing-web-service}books"/>
 *       &lt;/sequence>
 *     &lt;/restriction>
 *   &lt;/complexContent>
 * &lt;/complexType>
 * </pre>
 * 
 * 
 */
@XmlAccessorType(XmlAccessType.FIELD)
@XmlType(name = "", propOrder = {
    "books"
})
@XmlRootElement(name = "getBooksResponse")
public class GetBooksResponse {

    @XmlElement(required = true)
    protected Books books;

    /**
     * Obtient la valeur de la propriété books.
     * 
     * @return
     *     possible object is
     *     {@link Books }
     *     
     */
    public Books getBooks() {
        return books;
    }

    /**
     * Définit la valeur de la propriété books.
     * 
     * @param value
     *     allowed object is
     *     {@link Books }
     *     
     */
    public void setBooks(Books value) {
        this.books = value;
    }

}
