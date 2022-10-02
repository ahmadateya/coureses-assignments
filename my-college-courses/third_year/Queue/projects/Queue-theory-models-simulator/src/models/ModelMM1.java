
package models;


public class ModelMM1 {
    private long L, Lq;
    private double W, Wq;    
    
    public ModelMM1(double lambda, double mu){

        L = calcL(lambda, mu);// expected number of customers in the system
        Lq = calcLq(lambda, mu);// expected number of customers in the queue
        W = calcW(lambda, mu);// expected waiting time in the system (Wq + service time)
        Wq = calcWq(lambda, mu); // expected waiting time in the queue
        
    }

    public double getW() {
        return W;
    }

    public double getWq() {
        return Wq;
    }

    public long getL() {
        return L;
    }

    public long getLq() {
        return Lq;
    }

    // expected number of customers in the system
    private long calcL(double lambda, double mu) {
        return (long) Math.round(lambda / (mu - lambda));
    }
    // expected number of customers in the queue
    private long calcLq(double lambda, double mu) {
        return (long) Math.round((lambda * lambda) / (mu * (mu - lambda)));
    }
    // expected waiting time in the system (Wq + service time)
    private double calcW(double lambda, double mu) {
        return (1.0 / (mu - lambda));
    }
    // expected waiting time in the queue
    private double calcWq(double lambda, double mu) {
        return (lambda / (mu * (mu - lambda)));
    }

}
